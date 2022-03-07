package core

import (
	"context"
	"strings"

	"github.com/pedrocmart/canvas/internal/core/errors"
)

type CanvasGetRequest struct {
	ID string
}

type CanvasGetResponse struct {
	ID      string   `json:"CanvasID"`
	Drawing []string `json:"Drawing"`
}

func (h CanvasGetRequest) Validate() error {
	if h.ID == "" {
		return errors.BadRequest("empty ID")
	}
	return nil
}

func (c *canvasService) CanvasGet(ctx context.Context, req *CanvasGetRequest) (*CanvasGetResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, errors.Wrapf(err, "request validation")
	}

	canvas, err := c.store.CanvasGetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &CanvasGetResponse{
		ID:      canvas.ID,
		Drawing: strings.Split(canvas.Drawing, "\n"),
	}, nil
}
