//go:build postgres
// +build postgres

package pingers

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func init() {
	registerPinger("postgres", newPostgresConfiguration)
}

const defaultPostgresQuery = "SELECT 1"

func newPostgresConfiguration(connStr, query string) (*sql.DB, string, error) {
	if query == "" {
		query = defaultPostgresQuery
	}

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, query, err
	}

	return db, query, nil
}
