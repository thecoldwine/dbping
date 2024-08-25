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
)

func init() {
	p := pingers.ListPingers()

	if len(p) == 0 {
		panic("program had been compiled without any pingers, quit")
	}

	flags = flag.NewFlagSet("flags", flag.ExitOnError)
	flags.StringVar(&connectionString, "connection-string", "", "A connection string to the database, refer to the respective drivers")
	flags.StringVar(&dbtype, "dbtype", p[0], fmt.Sprintf("A database type to ping, supported databases in this build [%s]", strings.Join(p, ",")))
}

func main() {
	flags.Parse(os.Args)

	if len(os.Args) == 1 {
		flags.Usage()
	}

	_, err := pingers.Test(dbtype, connectionString, "")

	if err != nil {
		log.Fatalln(err)
	}

	// print results
}
