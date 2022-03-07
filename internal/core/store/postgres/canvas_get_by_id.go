package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pedrocmart/canvas/internal/core"
)

func (s *store) CanvasGetByID(ctx context.Context, id string) (*core.Canvas, error) {
	var canvas core.Canvas

	err := s.db.QueryRowContext(ctx, ` 
		SELECT
			id, 
			drawing, 
			created_at
		FROM canvas
		WHERE id = $1`, id).Scan(
		&canvas.ID,
		&canvas.Drawing,
		&canvas.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("canvas not found")
		}

		return nil, fmt.Errorf("CanvasGetByID: %s", err.Error())
	}
	return &canvas, nil
}
