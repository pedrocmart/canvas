package core_test

import (
	"context"
	"testing"
	"time"

	"github.com/matryer/is"
	"github.com/pedrocmart/canvas/internal/core"
	"github.com/pedrocmart/canvas/internal/core/errors"
	"github.com/pedrocmart/canvas/internal/core/log"
	"github.com/pedrocmart/canvas/internal/core/mocks"
)

func TestCanvasGet(t *testing.T) {
	now := time.Now()
	cases := []struct {
		it  string
		req *core.CanvasGetRequest

		//store
		canvasGetByIDStore   *core.Canvas
		canvasGetByIDError   error
		canvasGetByIDDrawing string

		expectedError  string
		expectedResult *core.CanvasGetResponse
	}{
		{
			it: "it create canvas",
			canvasGetByIDStore: &core.Canvas{
				ID:        "mock-id",
				Drawing:   `              .......`,
				CreatedAt: now,
			},
			canvasGetByIDDrawing: "mock-id",
			req: &core.CanvasGetRequest{
				ID: "mock-id",
			},
			expectedResult: &core.CanvasGetResponse{
				ID: "mock-id",
				Drawing: []string{
					"              .......",
				},
			},
		},
		{
			it: "it return error on CanvasGetByID",

			canvasGetByIDStore: &core.Canvas{},
			req: &core.CanvasGetRequest{
				ID: "mock-id",
			},
			canvasGetByIDDrawing: "mock-id",
			canvasGetByIDError:   errors.Wrap("mock-error"),
			expectedError:        "mock-error",
		},
		{
			it: "it return error on empty ID",

			canvasGetByIDStore: &core.Canvas{},
			req: &core.CanvasGetRequest{
				ID: "",
			},
			expectedError: "request validation: empty ID",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				CanvasGetByIDFunc: func(ctx context.Context, id string) (*core.Canvas, error) {
					checkIs.Equal(id, tc.canvasGetByIDDrawing)

					return tc.canvasGetByIDStore, tc.canvasGetByIDError
				},
			}

			service := core.New(
				&core.Config{
					Log: &log.Logger{},
				},
				store,
			)

			res, err := service.CanvasGet(context.Background(), tc.req)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
