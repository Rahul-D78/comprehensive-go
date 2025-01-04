package storage

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
)

type DBConfig struct {
	InventoryDatabaseAddr string
	InventoryDatabasePort int
	InventoryDatabaseUser string
	InventoryDatabasePass string
	InventoryDatabaseName string
}

// ConnectPostgres establish the connection to the database
// DSN(Data Source Name) format - postgres://username:password@address:port
// Sample connection string - postgres://postgres:mypassword@localhost:5432
// Connecting to user database - postgres://user:password@postgres:5432/user
func (r *DBConfig) ConnectPostgres(ctx context.Context, dbOpts ...DBConnOption) (*gorm.DB, error) {
	postgres_url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		r.InventoryDatabaseAddr, r.InventoryDatabaseUser,
		r.InventoryDatabasePass, r.InventoryDatabaseName, fmt.Sprint(r.InventoryDatabasePort),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  postgres_url,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to the DB", err)
		return nil, err
	}

	// Get the underlying sql.DB object
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("failed to get sql.DB from gorm.DB", err)
		return nil, err
	}

	for _, opt := range dbOpts {
		if err := opt(sqlDB); err != nil {
			return nil, fmt.Errorf("error applying option: %v", err)
		}
	}

	// retry every 5s up to 60s
	backoff := wait.Backoff{
		Steps:    12,
		Duration: 5 * time.Second,
		Factor:   1.0,
		Jitter:   0.1,
	}

	// check the connection
	err = retry.OnError(backoff, func(error) bool { return true }, func() error {
		if err := sqlDB.PingContext(ctx); err != nil {
			return fmt.Errorf("failed to ping the DB err: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to establish database connection err: %s", err)
	}

	fmt.Printf("Connected to db successfully %s", r.InventoryDatabaseName)
	return db, nil
}
