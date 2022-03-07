package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/matryer/is"
	"github.com/pedrocmart/canvas/internal/core"
	"github.com/pedrocmart/canvas/internal/core/errors"
	handlers "github.com/pedrocmart/canvas/internal/core/http"
	"github.com/pedrocmart/canvas/internal/core/log"
	"github.com/pedrocmart/canvas/internal/core/mocks"
	"github.com/pedrocmart/canvas/internal/core/testsuite"
)

func TestCanvasGet(t *testing.T) {
	cases := []struct {
		it string

		canvasGetResponse *core.CanvasGetResponse
		canvasGetError    error
		queryParams       string

		expectedError      string
		expectedResult     string
		expectedStatus     int
		mockReader         *testsuite.MockReader
		mockResponseWriter testsuite.MockResponseWriter
	}{
		{
			it: "it gets canvas",
			canvasGetResponse: &core.CanvasGetResponse{
				ID: "mock-canvas",
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
			expectedResult: `{"CanvasID":"mock-canvas","Drawing":["              .......","              .......","              .......","OOOOOOOO      .......","O      O      .......","O    XXXXX    .......","OOOOOXXXXX           ","     XXXXX           "]}`,
			expectedStatus: http.StatusOK,
		},
		{
			it: "it fails in service layer",
			canvasGetResponse: &core.CanvasGetResponse{
				ID:      "mock-canvas",
				Drawing: []string{},
			},
			canvasGetError: errors.Wrap("mock-error"),
			expectedResult: `{"error":{"message":"mock-error","status":500}}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				CanvasGetFunc: func(contextMoqParam context.Context, canvasGetRequest *core.CanvasGetRequest) (*core.CanvasGetResponse, error) {
					return tc.canvasGetResponse, tc.canvasGetError
				},
			}

			req, err := http.NewRequest("GET", "/"+tc.queryParams, nil)

			req = req.WithContext(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			h := handlers.NewHandler(&core.Config{
				Log: &log.Logger{},
			}, service)

			router := httprouter.New()
			rr := httptest.NewRecorder()

			router.GET("/", h.CanvasGet)
			router.ServeHTTP(rr, req)

			checkIs.Equal(rr.Body.String(), tc.expectedResult)
			checkIs.Equal(rr.Code, tc.expectedStatus)
		})
	}
}
