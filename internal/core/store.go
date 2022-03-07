package core

import (
	"context"
	"database/sql"
)

//go:generate moq -out ./mocks/store.go -pkg mocks  . Store
// Store interface
type Store interface {
	DB() *sql.DB

	CanvasCreate(context.Context, string) (string, error)
	CanvasGetByID(ctx context.Context, id string) (*Canvas, error)
}
