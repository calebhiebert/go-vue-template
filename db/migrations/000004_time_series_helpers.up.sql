CREATE OR REPLACE FUNCTION time_buckets(int_seconds integer, duration_seconds integer,
                                        start timestamp without time zone)
    RETURNS TABLE
            (
                bucket_start timestamp without time zone,
                bucket_end   timestamp without time zone
            )
AS
$$
SELECT to_timestamp(n) AS bucket_start, to_timestamp(n + int_seconds) AS bucket_end
FROM generate_series(
                 (EXTRACT(EPOCH FROM start)::int / int_seconds) * int_seconds,
                 EXTRACT(EPOCH FROM start)::int + duration_seconds,
                 int_seconds) n
$$
    LANGUAGE SQL;