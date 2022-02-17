package db

import (
	"database/sql"
	"github.com/XSAM/otelsql"
	lggr "github.com/datomar-labs-inc/FCT_Helpers_Go/logger"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func SetupDatabase() (*sql.DB, error) {
	driverName, err := otelsql.Register("postgres", semconv.DBSystemPostgreSQL.Value.AsString())
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(driverName, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(getDatabaseMaxConnections())

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDatabaseMaxConnections() int {
	envVal, err := strconv.Atoi(os.Getenv("DATABASE_MAX_CONNECTIONS"))
	if err != nil {

		if os.Getenv("DATABASE_MAX_CONNECTIONS") != "" {
			lggr.Get("get-db-max-conns").Warn("Invalid DATABASE_MAX_CONNECTIONS value, using default of 10", zap.String("value", os.Getenv("DATABASE_MAX_CONNECTIONS")))
		}

		return 10
	}

	return envVal
}
