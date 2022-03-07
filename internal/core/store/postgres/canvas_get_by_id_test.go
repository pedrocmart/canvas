package postgres_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/matryer/is"
	"github.com/pedrocmart/canvas/internal/core"
	"github.com/pedrocmart/canvas/internal/core/store/postgres"
)

func TestCanvasGetByID(t *testing.T) {
	now := time.Now()
	type row struct {
		id        string
		drawing   string
		createdAt time.Time
	}

	cases := []struct {
		it             string
		id             string
		r              *row
		expectedError  string
		sqlError       error
		rowError       error
		expectedResult *core.Canvas
	}{
		{
			it: "it returns *core.Canvas struct",
			id: "mock-id",
			r: &row{
				id:        "mock-id",
				drawing:   "",
				createdAt: now,
			},
			expectedResult: &core.Canvas{
				ID:        "mock-id",
				Drawing:   "",
				CreatedAt: now,
			},
		},
		{
			it:            "it fails to execute sql statement",
			id:            "mock-id",
			sqlError:      fmt.Errorf("mock-error"),
			expectedError: "CanvasGetByID: mock-error",
		},
		{
			it:            "it fails canvas not found",
			id:            "mock-id",
			sqlError:      sql.ErrNoRows,
			expectedError: "canvas not found",
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

			query := mock.ExpectQuery(regexp.QuoteMeta(`
			SELECT
				id, 
				drawing, 
				created_at
			FROM canvas
			WHERE id = $1`))
			switch {
			case tc.r != nil:
				mockrow := sqlmock.NewRows([]string{
					"id",
					"drawing",
					"created_at",
				}).AddRow(
					tc.r.id,
					tc.r.drawing,
					tc.r.createdAt,
				)
				query.WillReturnRows(mockrow)
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

			res, err := service.CanvasGetByID(context.Background(), tc.id)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
