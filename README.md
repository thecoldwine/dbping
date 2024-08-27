# dbping

`dbping` is a simple utility to check the network latency for the database server. It is useful
for comparing the network latency between different cloud providers or regions with the actual
database connection.

## Usage

```bash
# Usage of dbping:
#   -connection-string string
#         A connection string to the database, refer to the respective drivers
#   -dbtype string
#         A database type to ping, supported databases in this build [postgres] (default "postgres")
#   -pings int
#         A number of pings to the databases (default 1)
#   -query string
#         A query to execute for latency test. No sanity checks applied.

dbping --connection-string <connection-string> \
         --dbtype <dbtype> \
         --pings <pings> \
         --query <query>
```

## Build

Every pinger is defined with a tag. To build a support for a specific database, use the following command:

```bash
go build -tags <tag>
```

For example, to build a pinger for PostgreSQL:

```bash
go build -tags postgres
```

## Supported Databases

- PostgreSQL
- MySQL
- Azure SQL