package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/thecoldwine/dbping/pingers"
)

var (
	flags            *flag.FlagSet
	connectionString string
	dbtype           string
	query            string
	interval         time.Duration
	pings            int
)

func init() {
	p := pingers.ListPingers()

	if len(p) == 0 {
		panic("program had been compiled without any pingers, quit")
	}

	flag.StringVar(&connectionString, "connection-string", "", "A connection string to the database, refer to the respective drivers")
	flag.StringVar(&dbtype, "dbtype", p[0], fmt.Sprintf("A database type to ping, supported databases in this build [%s]", strings.Join(p, ",")))
	flag.StringVar(&query, "query", "", "A query to execute for latency test. No sanity checks applied.")
	flag.IntVar(&pings, "pings", 1, "A number of pings to the databases")
	flag.DurationVar(&interval, "interval", 0.0, "Interval between pings. Accepts duration format: 1s, 10ms, 1m and so on.")
}

func main() {
	if len(os.Args) == 1 {
		flag.Usage()
	}

	flag.Parse()

	log.Println("Testing a connection with driver:", dbtype)
	results, err := pingers.Test(dbtype, connectionString, query, pings, interval)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Total pings: %d. Min: %s; Max: %s; Average: %s.\n", results.Pings, results.Min, results.Max, results.Avg)

	if err != nil {
		log.Fatalln(err)
	}
}
