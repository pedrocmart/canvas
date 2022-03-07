package core_test

import (
	"context"
	"testing"

	"github.com/matryer/is"
	"github.com/pedrocmart/canvas/internal/core"
	"github.com/pedrocmart/canvas/internal/core/errors"
	"github.com/pedrocmart/canvas/internal/core/log"
	"github.com/pedrocmart/canvas/internal/core/mocks"
)

func TestCanvasCreate(t *testing.T) {
	cases := []struct {
		it  string
		req *core.CanvasCreateRequest

		//store
		canvasCreateStoreID      string
		canvasCreateStoreError   error
		canvasCreateStoreDrawing string

		expectedError  string
		expectedResult *core.CanvasCreateResponse
	}{
		{
			it:                  "it create canvas",
			canvasCreateStoreID: "mock-id",
			canvasCreateStoreDrawing: `              .......
              .......
              .......
OOOOOOOO      .......
O      O      .......
O    XXXXX    .......
OOOOOXXXXX           
     XXXXX           `,
			req: &core.CanvasCreateRequest{
				core.CanvasCreateRequestDetail{
					RectangleAt: [2]int{14, 0},
					Width:       7,
					Height:      6,
					Outline:     "none",
					Fill:        ".",
				},
				core.CanvasCreateRequestDetail{
					RectangleAt: [2]int{0, 3},
					Width:       8,
					Height:      4,
					Outline:     "O",
					Fill:        "none",
				},
				core.CanvasCreateRequestDetail{
					RectangleAt: [2]int{5, 5},
					Width:       5,
					Height:      3,
					Outline:     "X",
					Fill:        "X",
				},
			},
			expectedResult: &core.CanvasCreateResponse{
				Id: "mock-id",
				Drawing: []string{
					"              .......",
					"              .......",
					"              .......",
					"OOOOOOOO      .......",
					"O      O      .......",
					"O    XXXXX    .......",
					"OOOOOXXXXX           ",
					"     XXXXX           ",
				},
			},
		},
		{
			it: "it return error on CanvaCreate",

			canvasCreateStoreID: "mock-id",
			req: &core.CanvasCreateRequest{
				core.CanvasCreateRequestDetail{
					RectangleAt: [2]int{0, 0},
					Width:       1,
					Height:      1,
					Outline:     "none",
					Fill:        ".",
				},
			},
			canvasCreateStoreDrawing: ".",
			canvasCreateStoreError:   errors.Wrap("mock-error"),
			expectedError:            "failed to Create Canvas: mock-error",
		},
		{
			it: "it return error validate ASCII",

			canvasCreateStoreID: "mock-id",
			req: &core.CanvasCreateRequest{
				core.CanvasCreateRequestDetail{
					RectangleAt: [2]int{0, 0},
					Width:       1,
					Height:      1,
					Outline:     "¥",
					Fill:        "¥",
				},
			},
			expectedError: "request validation: We only accept ASCII characters :)",
		},
		{
			it: "it return error validate position lower than 0",

			canvasCreateStoreID: "mock-id",
			req: &core.CanvasCreateRequest{
				core.CanvasCreateRequestDetail{
					RectangleAt: [2]int{-1, 0},
					Width:       1,
					Height:      1,
					Outline:     ".",
					Fill:        ".",
				},
			},
			expectedError: "request validation: Position of rectangle cannot be lower than 0",
		},
		{
			it: "it return error validate size lower than canvas",

			canvasCreateStoreID: "mock-id",
			req: &core.CanvasCreateRequest{
				core.CanvasCreateRequestDetail{
					RectangleAt: [2]int{101, 0},
					Width:       1,
					Height:      1,
					Outline:     ".",
					Fill:        ".",
				},
			},
			expectedError: "request validation: The Canvas size is 100x100",
		},
		{
			it: "it return error validate width and height lower than 0",

			canvasCreateStoreID: "mock-id",
			req: &core.CanvasCreateRequest{
				core.CanvasCreateRequestDetail{
					RectangleAt: [2]int{0, 0},
					Width:       -1,
					Height:      1,
					Outline:     ".",
					Fill:        ".",
				},
			},
			expectedError: "request validation: Width and Height must be greater than 0",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				CanvasCreateFunc: func(contextMoqParam context.Context, s string) (string, error) {
					checkIs.Equal(s, tc.canvasCreateStoreDrawing)
					return tc.canvasCreateStoreID, tc.canvasCreateStoreError
				},
			}

			service := core.New(
				&core.Config{
					Log: &log.Logger{},
				},
				store,
			)

			res, err := service.CanvasCreate(context.Background(), tc.req)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
