package main

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type RequestStats struct {
	BucketStart  time.Time `json:"bucket_start"`
	Max          int       `json:"max"`
	Average      float64   `json:"average"`
	Percentile50 int       `json:"percentile_50"`
	Percentile75 int       `json:"percentile_75"`
	Percentile90 int       `json:"percentile_90"`
	Percentile95 int       `json:"percentile_95"`
	Percentile99 int       `json:"percentile_99"`
	Count        int       `json:"count"`
}

func GetAdminDashboardStats(c *gin.Context) {
	userCount, err := models.Users().CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	newUsersMonthly, err := models.Users(qm.
		Where("created_at BETWEEN (now() at time zone 'utc' - interval '1 month') AND (now() at time zone 'utc')")).
		CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	newUsersWeekly, err := models.Users(qm.
		Where("created_at BETWEEN (now() at time zone 'utc' - interval '1 week') AND (now() at time zone 'utc')")).
		CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	reqStats, err := GetRequestStats(c.Request.Context(), 5*time.Minute, 24*time.Hour, time.Now().Add(-24*time.Hour))
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_users":       userCount,
		"monthly_new_users": newUsersMonthly,
		"weekly_new_users":  newUsersWeekly,
		"request_stats":     reqStats,
	})
}

func GetRequestStats(ctx context.Context, interval time.Duration, duration time.Duration, start time.Time) ([]RequestStats, error) {
	var res []RequestStats

	err := queries.Raw(fmt.Sprintf(`
	SELECT i.bucket_start,
       COALESCE(max(processing_duration), 0) AS max,
       COALESCE(avg(processing_duration), 0) AS avg,
       percentile_disc(0.50) WITHIN GROUP ( ORDER BY processing_duration ) AS percentile_50,
       percentile_disc(0.75) WITHIN GROUP ( ORDER BY processing_duration ) AS percentile_75,
       percentile_disc(0.90) WITHIN GROUP ( ORDER BY processing_duration ) AS percentile_90,
       percentile_disc(0.95) WITHIN GROUP ( ORDER BY processing_duration ) AS percentile_95,
       percentile_disc(0.99) WITHIN GROUP ( ORDER BY processing_duration ) AS percentile_99,
       count(*) AS count
    FROM access_logs a
        RIGHT JOIN time_buckets(%d, %d, to_timestamp(%d)::timestamptz at time zone 'utc') i
            ON a.created_at >= i.bucket_start
            AND a.created_at < i.bucket_end
    GROUP BY i.bucket_start
    ORDER BY i.bucket_start DESC;`, int(math.Round(interval.Seconds())), int(math.Round(duration.Seconds())), start.Unix())).
		BindG(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
