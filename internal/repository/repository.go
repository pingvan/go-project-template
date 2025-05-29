package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(ctx context.Context, connString string) (*Repository, error) {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("could not open databse connection: %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("could not establish connection with databse: %w", err)
	}
	return &Repository{db: db}, nil
}
