//go:build mysql
// +build mysql

package pingers

func init() {
	registerPinger("mysql", newPostgresPinger)
}

type mysqlPinger struct {
}

func newmysqlPinger(connStr, sql string) pinger {
	return &mysqlPinger{}
}

func (r *mysqlPinger) Connect() error {
	return nil
}

func (r *mysqlPinger) Ping() error {
	return nil
}

func (r *mysqlPinger) Close() {

}
