package store

import (
	"database/sql"

	"github.com/pedrocmart/canvas/internal/core"
	"github.com/pedrocmart/canvas/internal/core/store/postgres"
)

// NewStore create new store
func NewStore(db *sql.DB) core.Store {
	return postgres.New(db)
}
