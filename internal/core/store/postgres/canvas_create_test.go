package postgres_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/matryer/is"
	"github.com/pedrocmart/canvas/internal/core/store/postgres"
)

func TestCanvasCreate(t *testing.T) {
	now := time.Now()
	type row struct {
		ID string
	}
	cases := []struct {
		it            string
		drawing       string
		sqlError      error
		created_at    time.Time
		r             *row
		rowError      error
		expectedError string
	}{
		{
			it: "it inserts canvas",
			r: &row{
				ID: "mock-id",
			},
			drawing:    "teste1",
			created_at: now,
		},
		{
			it:            "it fails to execute sql statement",
			drawing:       "teste1",
			created_at:    now,
			sqlError:      fmt.Errorf("mock-error"),
			expectedError: "create canvas: mock-error",
		},
	}
	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			query := mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO canvas (drawing) VALUES ($1) RETURNING id`))

			switch {
			case tc.r != nil:
				mockRow := sqlmock.NewRows([]string{
					"id",
				}).AddRow(
					tc.r.ID,
				)
				query.WillReturnRows(mockRow)
			case tc.rowError != nil:
				mockRow := sqlmock.NewRows([]string{
					"id",
				}).AddRow(nil).RowError(1, tc.rowError)
				query.WillReturnRows(mockRow)
			default:
				query.WillReturnRows(sqlmock.NewRows([]string{
					"id",
				}).AddRow(""))
				query.WillReturnError(tc.sqlError)
			}

			service := postgres.New(db)

			_, err = service.CanvasCreate(context.Background(), tc.drawing)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
		})
	}
}
