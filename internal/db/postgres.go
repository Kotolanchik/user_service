package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"time"
	loggingCodes "user-service/internal/types/logging"
	logging "user-service/pkg/logger"
)

var logger = logging.New(logging.GetDefaultConfig())

const (
	maxConnectionAttempts = 5
	retryInterval         = time.Second
)

func CreateConnectin(conn string) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	for attempt := 1; attempt <= maxConnectionAttempts; attempt++ {
		db, err = tryCreateConn(conn)
		if err == nil {
			return db, nil
		}
		logger.Error(loggingCodes.CodeWarmAttemptConnection, fmt.Sprintf("Attempt %d to connect to PostgreSQL failed: ", attempt), err)
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("failed to connect to PostgreSQL after %d attempts", maxConnectionAttempts)
}

func tryCreateConn(conn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", conn) // sqlx.ConnectContext
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			logger.Error(loggingCodes.CodeErrorDBConnectionFailed, "PostgreSQL Error code and msg "+pqErr.Code.Name(), pqErr.Message)
		}

		return nil, err
	}

	return db, nil
}
