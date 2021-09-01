package db

import (
	"context"
	"database/sql"
	"os"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/luna-duclos/instrumentedsql"
	isqlot "github.com/luna-duclos/instrumentedsql/opentracing"
)

func SetupDatabase() (*sql.DB, error) {
	logger := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
		// fmt.Printf("%s %v\n", msg, keyvals)
	})

	sql.Register("instrumented-pgsql", instrumentedsql.WrapDriver(pq.Driver{}, instrumentedsql.WithTracer(isqlot.NewTracer(true)), instrumentedsql.WithLogger(logger)))

	db, err := sql.Open("instrumented-pgsql", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
