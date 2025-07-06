package database

import (
	"context"
	"fmt"
	"golang-agnostic-template/src/pkg/config"

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
	db, err := surrealdb.New(fmt.Sprintf("ws://%s:%d", config.Params.DBHost, config.Params.DBPort))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SurrealDB: %w", err)
	}
	authData := &surrealdb.Auth{
		Username: config.Params.DBUser,
		Password: config.Params.DBPassword,
	}
	_, err = db.SignIn(authData)
	if err != nil {
		return nil, fmt.Errorf("failed to sign in to SurrealDB: %w", err)
	}

	if err := db.Use(config.Params.DBNamespace, config.Params.DBDatabase); err != nil {
		return nil, fmt.Errorf("failed to use database: %w", err)
	}
	s.db = db
	return db, nil
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
