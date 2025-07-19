package database

import (
	"context"
	"fmt"
	"golang-agnostic-template/src/pkg/config"
	"time"

	"github.com/surrealdb/surrealdb.go"
)

type IDatabaseConnection interface {
	Connect(ctx context.Context) (*surrealdb.DB, error)
	DB() *surrealdb.DB
	Close() error
	Use(namespace, database string) error
}

type SurrealDBConnection struct {
	db *surrealdb.DB
}

func NewSurrealDBConnection() IDatabaseConnection {
	return &SurrealDBConnection{}
}

func (s *SurrealDBConnection) Connect(ctx context.Context) (*surrealdb.DB, error) {
	const maxAttempts = 3
	const retryDelay = 3 * time.Second

	dsn := fmt.Sprintf("ws://%s:%d", config.Params.DBHost, config.Params.DBPort)
	var db *surrealdb.DB
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = surrealdb.New(dsn)
		if err == nil {
			authData := &surrealdb.Auth{
				Username: config.Params.DBUser,
				Password: config.Params.DBPassword,
			}
			_, err = db.SignIn(authData)
			if err == nil {
				err = db.Use(config.Params.DBNamespace, config.Params.DBDatabase)
				if err == nil {
					s.db = db
					return db, nil
				}
			}
		}
		fmt.Printf("Attempt %d/%d failed: %v\n", attempt, maxAttempts, err)

		if attempt < maxAttempts {
			select {
			case <-time.After(retryDelay):
			case <-ctx.Done():
				return nil, fmt.Errorf("connection attempt cancelled: %w", ctx.Err())
			}
		}
	}
	return nil, fmt.Errorf("failed to connect to SurrealDB after %d attempts: %w", maxAttempts, err)
}

func (s *SurrealDBConnection) DB() *surrealdb.DB {
	return s.db
}

func (s *SurrealDBConnection) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

func (s *SurrealDBConnection) Use(namespace, database string) error {
	return s.db.Use(namespace, database)
}
