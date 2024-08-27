//go:build postgres
// +build postgres

package pingers

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func init() {
	registerPinger("postgres", newPostgresPinger)
}

const defaultPgSql = "SELECT 1"

type postgresPinger struct {
	conn    *pgx.Conn
	connStr string
	sql     string
}

func newPostgresPinger(connStr, sql string) pinger {
	if sql == "" {
		sql = defaultPgSql
	}

	return &postgresPinger{
		connStr: connStr,
		sql:     sql,
	}
}

func (r *postgresPinger) Connect() error {
	log.Println("connecting with", r.connStr)

	conn, err := pgx.Connect(context.TODO(), r.connStr)
	if err != nil {
		return err
	}
	r.conn = conn

	return nil
}

func (r *postgresPinger) Ping() error {
	row := r.conn.QueryRow(context.TODO(), r.sql)
	var discard int

	if err := row.Scan(&discard); err != nil {
		return err
	}

	return nil
}

func (r *postgresPinger) Close() {
	if r.conn != nil && !r.conn.IsClosed() {
		r.conn.Close(context.TODO())
	}
}
