//go:build postgres
// +build postgres

package pingers

func init() {
	registerPinger("postgres", newPostgresPinger)
}

type postgresPinger struct {
}

func newPostgresPinger(connStr, sql string) pinger {
	return &postgresPinger{}
}

func (r *postgresPinger) Connect() error {
	return nil
}

func (r *postgresPinger) Ping() error {
	return nil
}

func (r *postgresPinger) Close() {

}
