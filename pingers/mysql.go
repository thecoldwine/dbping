//go:build mysql
// +build mysql

package pingers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	registerPinger("mysql", newMysqlConfiguration)
}

const defaultMysqlQuery = "SELECT 1"

func newMysqlConfiguration(connStr, query string) (*sql.DB, string, error) {
	if query == "" {
		query = defaultMysqlQuery
	}

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, query, err
	}

	return db, query, nil
}
