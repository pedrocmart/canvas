package postgres

import (
	"context"

	"github.com/lib/pq"
	"github.com/pedrocmart/canvas/internal/core/errors"
)

func (s *store) CanvasCreate(ctx context.Context, drawing string) (string, error) {
	var id string
	err := s.db.QueryRowContext(ctx, `
		INSERT INTO canvas (drawing) VALUES ($1) RETURNING id`,
		drawing,
	).Scan(&id)
	if err != nil {
		p, ok := err.(*pq.Error)
		if ok {
			err = errors.Wrapf(err, " database error: %s", p.Code.Class().Name())
		}

		return "", errors.Wrapf(err, "create canvas")
	}

	return id, nil
}
