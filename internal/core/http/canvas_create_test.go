package http_test

import (
	"bytes"
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

func TestCanvasCreate(t *testing.T) {
	mockReader := testsuite.MockReader(0)

	cases := []struct {
		it                   string
		canvasCreateResponse *core.CanvasCreateResponse
		canvasSetError       error
		queryParams          string

		expectedError      string
		expectedResult     string
		expectedStatus     int
		mockReader         *testsuite.MockReader
		mockResponseWriter testsuite.MockResponseWriter
		requestBody        string
	}{
		{
			it:                   "it create canvas",
			expectedResult:       `{"CanvasID":"mock-canvas-id","Drawing":["."]}`,
			expectedStatus:       http.StatusOK,
			canvasCreateResponse: &core.CanvasCreateResponse{Id: "mock-canvas-id", Drawing: []string{"."}},
			requestBody: `[{
				"RectangleAt" : [0,0],
				"Width": 1,
				"Height": 1,
				"Outline": "none",
				"Fill": "."
			}
			]`,
		},
		{
			it:                   "it fails to read the request body",
			expectedResult:       `{"error":{"message":"test error","status":500}}`,
			expectedStatus:       http.StatusInternalServerError,
			canvasCreateResponse: &core.CanvasCreateResponse{Id: "mock-canvas-id"},
			requestBody: `[{
				"RectangleAt" : [14,0],
				"Width": 7,
				"Height": 6,
				"Outline": "none",
				"Fill": "."
			},
			{
				"RectangleAt" : [0,3],
				"Width": 8,
				"Height": 4,
				"Outline": "O",
				"Fill": "none"
			},
			{
				"RectangleAt" : [5,5],
				"Width": 5,
				"Height": 3,
				"Outline": "X",
				"Fill": "X"
			}
			]`,
			mockReader: &mockReader,
		},
		{
			it:                   "it fails to parse invalid json body",
			expectedResult:       `{"error":{"message":"invalid character ':' after top-level value","status":500}}`,
			expectedStatus:       http.StatusInternalServerError,
			canvasCreateResponse: &core.CanvasCreateResponse{Id: "mock-canvas-id"},
			requestBody: `
				"RectangleAt" : [14,0],
				"Width": 7,
				"Height": 6,
				"Outline": "none",
				"Fill": "."
			},
			{
				"RectangleAt" : [0,3],
				"Width": 8,
				"Height": 4,
				"Outline": "O",
				"Fill": "none"
			},
			{
				"RectangleAt" : [5,5],
				"Width": 5,
				"Height": 3,
				"Outline": "X",
				"Fill": "X"
			}
			]`,
		},
		{
			it:                   "it create canvas",
			expectedResult:       `{"error":{"message":"mock-error","status":500}}`,
			expectedStatus:       http.StatusInternalServerError,
			canvasCreateResponse: &core.CanvasCreateResponse{Id: "mock-canvas-id"},
			requestBody: `[{
				"RectangleAt" : [14,0],
				"Width": 7,
				"Height": 6,
				"Outline": "none",
				"Fill": "."
			},
			{
				"RectangleAt" : [0,3],
				"Width": 8,
				"Height": 4,
				"Outline": "O",
				"Fill": "none"
			},
			{
				"RectangleAt" : [5,5],
				"Width": 5,
				"Height": 3,
				"Outline": "X",
				"Fill": "X"
			}
			]`,
			canvasSetError: errors.Wrap("mock-error"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				CanvasCreateFunc: func(contextMoqParam context.Context, canvasCreateRequest *core.CanvasCreateRequest) (*core.CanvasCreateResponse, error) {
					return tc.canvasCreateResponse, tc.canvasSetError
				},
			}

			var req *http.Request
			var err error
			if tc.mockReader == nil {
				req, err = http.NewRequest("POST", "/", bytes.NewReader([]byte(tc.requestBody)))
			} else {
				req, err = http.NewRequest("POST", "/", tc.mockReader)
			}

			req = req.WithContext(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			h := handlers.NewHandler(&core.Config{
				Log: &log.Logger{},
			}, service)

			router := httprouter.New()
			rr := httptest.NewRecorder()

			router.POST("/", h.CanvasCreate)
			router.ServeHTTP(rr, req)

			checkIs.Equal(rr.Body.String(), tc.expectedResult)
			checkIs.Equal(rr.Code, tc.expectedStatus)
		})
	}
}
