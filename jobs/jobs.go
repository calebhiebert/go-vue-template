package jobs

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/calebhiebert/go-vue-template/util"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

const CleanFailedJobs = false
const CleanCompletedJobs = true
const CleanCompletedJobsAfter = 5 * time.Minute

type EnqueueJobOpts struct {
	Type       *string
	Priority   *int
	Source     *string
	Data       interface{}
	RunAt      *time.Time
	MaxRetries int
}

type FailureData struct {
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type worker struct {
	runner  JobRunnerFunc
	running bool
	jobType string
}

var Workers map[string]*worker

type JobRunnerFunc func(ctx context.Context, j *models.Job) error

func grabJob(ctx context.Context, jobType string) (*models.Job, error) {
	var j models.Job

	err := queries.RawG(`SELECT * FROM get_job($1)`, jobType).BindG(ctx, &j)
	if err != nil {
		return nil, err
	}

	return &j, nil
}

func InitJobs() error {
	AddWorker("job-cleanup", 1, func(ctx context.Context, j *models.Job) error {
		if CleanFailedJobs {
			count, err := models.Jobs(qm.Where("status = 'failed' AND (max_retries = 0 OR retry_count = max_retries)")).
				DeleteAllG(ctx)
			if err != nil {
				fmt.Println("Error cleaning up failed jobs")
			} else {
				fmt.Printf("Cleaned up %d failed jobs\n", count)
			}
		}
		
		if CleanCompletedJobs {
			count, err := models.Jobs(qm.Where("status = 'done' AND finished_at > ?", time.Now().Add(CleanCompletedJobsAfter))).
				DeleteAllG(ctx)
			if err != nil {
				fmt.Println("Error cleaning up finished jobs")
			} else {
				fmt.Printf("Cleaned up %d finished jobs\n", count)
			}
		}
		
		_, err := EnqueueJob(ctx, &EnqueueJobOpts{
			Type:       util.StrPtr("job-cleanup"),
			Source:     util.StrPtr("auto-job"),
			RunAt:      util.TimePtr(time.Now().UTC().Add(1 * time.Hour)),
			MaxRetries: 2,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if exists, err := models.Jobs(qm.Where("type = 'job-cleanup' AND status = 'fresh'")).ExistsG(context.Background()); !exists && err == nil {
		_, err := EnqueueJob(context.Background(), &EnqueueJobOpts{
			Type:       util.StrPtr("job-cleanup"),
			Source:     util.StrPtr("auto-job"),
			RunAt:      util.TimePtr(time.Now().UTC().Add(1 * time.Hour)),
			MaxRetries: 2,
		})
		if err != nil {
			return err
		}
	} else if err != nil {
		panic(err)
	}

	return nil
}

func AddWorker(jobType string, threads int, f JobRunnerFunc) {
	if existingWorker, ok := Workers[jobType]; ok {
		existingWorker.running = false
	}

	if Workers == nil {
		Workers = make(map[string]*worker)
	}

	Workers[jobType] = &worker{
		runner:  f,
		running: true,
		jobType: jobType,
	}

	workCh := make(chan *models.Job)

	for i := 0; i < threads; i++ {
		go func(w *worker, c chan *models.Job) {
			fmt.Println("Worker started! ", w.jobType)

			// Get a job from the channel
			for j := range c {
				// Process the job
				err := w.runner(context.Background(), j)
				if err != nil {
					var fd []FailureData

					if j.FailureData.Valid {
						err := j.FailureData.Unmarshal(&fd)
						if err != nil {
							fmt.Println("Failed to unmarshal failure data", err)
						}
					}

					fd = append(fd, FailureData{
						Message: err.Error(),
						Details: err,
					})

					err := j.FailureData.Marshal(fd)
					if err != nil {
						fmt.Println("Failed to marshal failure data", err)
					}

					j.Status = "failed"

					if j.MaxRetries > 0 {
						j.RunAt = null.TimeFrom(time.Now().Add(time.Duration(j.RetryCount.Int) * time.Second).UTC())
					}

					_, err = j.UpdateG(context.Background(), boil.Infer())
					if err != nil {
						fmt.Println("Failed to update failed job")
					}
				} else {
					j.Status = "done"
					j.FinishedAt = null.TimeFrom(time.Now().UTC())

					_, err := j.UpdateG(context.Background(), boil.Infer())
					if err != nil {
						fmt.Println("Failed to update job as success")
					}
				}
			}
		}(Workers[jobType], workCh)
	}

	go func(w *worker, c chan *models.Job) {
		for w.running {
			j, err := grabJob(context.Background(), w.jobType)
			if err == sql.ErrNoRows {
				time.Sleep(5 * time.Second)
			} else if err != nil {
				fmt.Println("Failed to grab job", err)
			} else {
				c <- j
			}
		}

		close(c)
	}(Workers[jobType], workCh)
}

func EnqueueJob(ctx context.Context, opts *EnqueueJobOpts) (*models.Job, error) {
	var j = &models.Job{}

	j.ID = uuid.Must(uuid.NewV4()).String()

	if opts.Type != nil {
		j.Type = null.StringFromPtr(opts.Type)
	}

	if opts.Priority != nil {
		j.Priority = *opts.Priority
	}

	if opts.Source != nil {
		j.Source = null.StringFromPtr(opts.Source)
	}

	err := j.Data.Marshal(opts.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	if opts.RunAt != nil {
		j.RunAt = null.TimeFromPtr(opts.RunAt)
	}

	j.MaxRetries = opts.MaxRetries

	err = j.InsertG(ctx, boil.Infer())
	if err != nil {
		return nil, err
	}

	return j, nil
}
