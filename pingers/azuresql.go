//go:build azuresql
// +build azuresql

package pingers

import (
	"database/sql"

	"github.com/microsoft/go-mssqldb/azuread"
)

func init() {
	registerPinger("azuresql", newAzureSqlConfiguration)
}

const defaultAzureSqlQuery = "SELECT 1"

func newAzureSqlConfiguration(connStr, query string) (*sql.DB, string, error) {
	if query == "" {
		query = defaultAzureSqlQuery
	}

	db, err := sql.Open(azuread.DriverName, connStr)
	if err != nil {
		return nil, query, err
	}

	return db, query, nil
}
