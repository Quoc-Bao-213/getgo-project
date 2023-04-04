package pg

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/pkg/errors"
	_ "github.com/lib/pq"
)

func Connect(_ context.Context) (*sql.DB, error) {
	// `database` host is container which already defined in docker compose
	conn, err := sql.Open("postgres", "host=db user=postgres password=postgres dbname=getgo-pg port=5432 sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

		return nil, errors.WithStack(fmt.Errorf("connect DB failed. Err: %w", err))
	}

	if err = conn.Ping(); err != nil {

		return nil, errors.WithStack(err)
	}

	fmt.Fprintln(os.Stderr, "connect to DB successfully")

	return conn, nil
}
