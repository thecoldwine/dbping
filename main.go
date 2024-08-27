package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/thecoldwine/dbping/pingers"
)

var (
	flags            *flag.FlagSet
	connectionString string
	dbtype           string
	query            string
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
}

func main() {
	if len(os.Args) == 1 {
		flag.Usage()
	}

	flag.Parse()

	results, err := pingers.Test(dbtype, connectionString, query, pings)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Total pings: %d. Min: %s; Max: %s; Average: %s.\n", results.Pings, results.Min, results.Max, results.Avg)

	if err != nil {
		log.Fatalln(err)
	}
}
