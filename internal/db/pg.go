package db

import (
	"database/sql"
	"room-reservation/config"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

func NewDBConnection(config config.PostgresConfig, logger *zap.Logger) (*sql.DB, error) {
	//Connect to pg
	db, err := sql.Open("pgx", config.ConnectionString)
	logger.Debug("Connecting to db", zap.String("connection_string", config.ConnectionString))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(29 * time.Minute) //Azure's default is 30 mins, so we set it to 29 mins to be safe.

	// Pings the database.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
