package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lib/pq"
	"github.com/pedrocmart/canvas/internal/core"
	"github.com/pedrocmart/canvas/internal/core/errors"
	"github.com/pedrocmart/canvas/internal/core/store"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

const (
	retryTimeout time.Duration = 5 * time.Second
	postgresPort int           = 5432
)

func NewStore(c *core.Config) (core.Store, error) {
	host := os.Getenv("CANVAS_DB_HOST")
	user := os.Getenv("CANVAS_DB_USER")
	password := os.Getenv("CANVAS_DB_PASSWORD")
	dbname := os.Getenv("CANVAS_DB_NAME")

	if host == "" {
		return nil, errors.Wrap("database host is empty")
	}

	if user == "" {
		return nil, errors.Wrap("database user is empty")
	}

	if password == "" {
		return nil, errors.Wrap("database password is empty")
	}

	if dbname == "" {
		return nil, errors.Wrap("database name is empty")
	}

	sqltrace.Register("postgres", pq.Driver{})

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, postgresPort, user, password, dbname)

	db, err := retryConn(psqlInfo)
	if err != nil {
		return nil, errors.Wrapf(err, "store connection error")
	}

	return store.NewStore(db), nil
}

func retryConn(psqlInfo string) (*sql.DB, error) {
	for i := 0; i <= 3; i++ {
		db, err := sqltrace.Open("postgres", psqlInfo, sqltrace.WithServiceName(core.ServiceName))
		if err != nil {
			log.Println("postgres connection error:", err)
			time.Sleep(retryTimeout)

			continue
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		time.Sleep(retryTimeout)
	}

	return nil, errors.Wrap("database connection retry exceded")
}
