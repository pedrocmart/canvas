package postgres

import (
	"context"
	"database/sql"
	"strings"

	"github.com/pedrocmart/canvas/internal/core"
	"github.com/pedrocmart/canvas/internal/core/errors"
)

// New create new postgres store
func New(db *sql.DB) core.Store {
	s := &store{
		db: db,
	}

	return s
}

type store struct {
	db *sql.DB
}

func (s *store) DB() *sql.DB {
	return s.db
}

// Ping test db connection
func (s *store) Ping(ctx context.Context) error {
	if err := s.db.PingContext(ctx); err != nil {
		return errors.Wrapf(err, "ping postgres")
	}

	return nil
}

// SliceToString prepare string slice to string query
// pq.Array(&someSlice) this shit return double quotes, not single for postgres query
// check with log.Print(pq.Array(&someSlice).Value())
// use this for concatante slice with single quotes
func SliceToString(s []string) string {
	return "'" + strings.Join(s, "','") + "'"
}
