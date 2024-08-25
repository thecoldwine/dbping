package pingers

// internal type for the pinger
// normally we just need a connection string and sql statement in _some_ cases
// if the sql statement is an empty string the default one will be used
type pingerFactory func(connStr, sqlStatement string) pinger

var pingers map[string]pingerFactory

func registerPinger(dbtype string, factory pingerFactory) {
	if pingers == nil {
		pingers = make(map[string]pingerFactory, 0)
	}

	pingers[dbtype] = factory
}

type Results struct {
	Attempts     int
	TotalLatency float64
	Errors       []error
}

type pinger interface {
	Connect() error
	Ping() error
	Close()
}

func ListPingers() []string {
	results := make([]string, 0, len(pingers))

	for k := range pingers {
		results = append(results, k)
	}

	return results
}

func Test(dbtype, connectionString, sql string) (*Results, error) {
	pinger := pingers[dbtype](connectionString, sql)

	err := pinger.Connect()
	if err != nil {
		return nil, err
	}

	defer pinger.Close()

	return nil, nil
}
