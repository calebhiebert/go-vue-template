CREATE TABLE IF NOT EXISTS jobs
(
    id           UUID    NOT NULL PRIMARY KEY,

    type         VARCHAR,
    priority     INTEGER NOT NULL            DEFAULT 1,
    source       VARCHAR,
    data         JSONB,
    run_at       TIMESTAMP WITHOUT TIME ZONE, -- Null run_at means run job immediately

    failure_data JSONB,
    max_retries  INTEGER NOT NULL            DEFAULT 0,
    retry_count  INTEGER,
    status       VARCHAR NOT NULL            DEFAULT 'fresh',
    finished_at  TIMESTAMP WITHOUT TIME ZONE,

    created_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc')
);

create OR REPLACE function get_job(jobtype VARCHAR) returns setof jobs
    language plpgsql
as
$$
declare
    j jobs;
begin
    SELECT *
    INTO j
    FROM jobs
    WHERE (now() > jobs.run_at
        OR run_at IS NULL)
      AND (status = 'fresh' OR
           (status = 'failed' AND max_retries > 0 AND (retry_count IS NULL OR retry_count < jobs.max_retries)))
      AND CASE
              WHEN jobType = 'nil'
                  THEN type IS NULL
              ELSE type = jobType
        END
    ORDER BY priority DESC
        FOR UPDATE SKIP LOCKED
    LIMIT 1;

    IF j.id IS NULL THEN
        return;
    end if;

    IF j.status = 'failed' THEN
        IF j.retry_count IS NULL THEN
            j.retry_count = 1;
        ELSE
            j.retry_count = j.retry_count + 1;
        end if;
    end if;

    j.status = 'grabbed';
    UPDATE jobs
    SET status      = j.status,
        retry_count = j.retry_count
    WHERE id = j.id;

    return next j;
end;
$$;