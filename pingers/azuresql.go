//go:build azuresql
// +build azuresql

package pingers

func init() {
	registerPinger("azuresql", newAzureSqlPinger)
}

type azureSqlPinger struct {
}

func newAzureSqlPinger(connStr, sql string) pinger {
	return &azureSqlPinger{}
}

func (r *azureSqlPinger) Connect() error {
	return nil
}

func (r *azureSqlPinger) Ping() error {
	return nil
}

func (r *azureSqlPinger) Close() {

}
