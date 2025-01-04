package storage

import (
	"database/sql"
	"fmt"
	"time"
)

// DBConnOption defines a function type for configuring the database connection.
type DBConnOption func(*sql.DB) error

func WithTimeout(timeout time.Duration) DBConnOption {
	return func(c *sql.DB) error {
		c.SetConnMaxIdleTime(timeout)
		return nil
	}
}

func WithMaxConnections(maxOpenDbConn int) DBConnOption {
	return func(c *sql.DB) error {
		if maxOpenDbConn <= 0 {
			return fmt.Errorf("maxOpenConns must be positive")
		}
		c.SetMaxIdleConns(maxOpenDbConn)
		return nil
	}
}
