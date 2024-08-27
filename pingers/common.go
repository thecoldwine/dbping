package pingers

import (
	"database/sql"
	"log"
	"math"
	"time"
)

// internal type for the pinger
// normally we just need a connection string and sql statement in _some_ cases
// if the sql statement is an empty string the default one will be used
type pingerFactory func(connStr, query string) (*sql.DB, string, error)

var pingers map[string]pingerFactory

func registerPinger(dbtype string, factory pingerFactory) {
	if pingers == nil {
		pingers = make(map[string]pingerFactory, 0)
	}

	pingers[dbtype] = factory
}

type Results struct {
	Pings  int
	Avg    time.Duration
	Min    time.Duration
	Max    time.Duration
	Errors []error
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

func Test(dbtype, connectionString, sql string, pings int) (*Results, error) {
	moment := time.Now()

	db, query, err := pingers[dbtype](connectionString, sql)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	defer db.Close()

	log.Println("Connected in", time.Since(moment))

	errs := make([]error, 0)

	min := math.Inf(1)
	max := math.Inf(-1)
	sum := 0.0
	for i := range pings {
		moment = time.Now()

		_, err := db.Exec(query)
		if err != nil {
			errs = append(errs, err)
			log.Println("error while ping", i, ":", err)
		}

		elapsed := time.Since(moment)
		min = math.Min(min, float64(elapsed))
		max = math.Max(max, float64(elapsed))

		sum += float64(elapsed)
	}

	results := &Results{
		Pings:  pings,
		Avg:    time.Duration(sum / float64(pings)),
		Min:    time.Duration(min),
		Max:    time.Duration(max),
		Errors: errs,
	}

	return results, nil
}
