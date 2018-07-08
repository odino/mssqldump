package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs"
	_ "github.com/denisenkom/go-mssqldb"
	flags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"os"
)

const version string = "1.1.0"

var opts struct {
	Query       string `short:"q" long:"query" required:"true" description:"Query to run to extract data"`
	Format      string `short:"f" long:"format" required:"false" description:"Output format (default: TSV)"`
	ColumnNames bool   `short:"c" long:"column-names" required:"false" description:"Whether to include column names as first line of the output file"`
}

func main() {
	setupLogging()
	_, err := flags.Parse(&opts)
	handleError(err)
	log.Debugf("options: %v", opts)

	if opts.Query == "v" || opts.Query == "version" {
		fmt.Println(version)
		os.Exit(0)
	}

	c := getConnection()
	log.Debugf("connection: %s", c)

	rows := query(c, opts.Query)
	output(rows, opts.Format, opts.ColumnNames)
}

// Creates a new writer, defaulting to tabs
// as separator.
func newTsvWriter() *csv.Writer {
	w := csv.NewWriter(os.Stdout)
	w.Comma = '\t'

	return w
}

// Utility function to setup logging.
// If no log level is specified from
// the environment, we default to ERROR,
// so that the resulting TSV doesn't get mixed
// up with logging statements unless an error
// occurred.
func setupLogging() {
	l := os.Getenv("LOG_LEVEL")
	if l == "" {
		l = "ERROR"
	}
	level, err := log.ParseLevel(l)
	handleError(err)
	log.SetLevel(level)
}

// Receives a list of rows and iterates through them,
// while printing them to the stdout
func output(rows *sql.Rows, format string, includeColumns bool) {
	defer rows.Close()

	if format == "json" {
		outputJson(rows)
	} else {
		outputTsv(rows, includeColumns)
	}

	log.Debug("done")
}

func outputTsv(rows *sql.Rows, includeColumns bool) {
	w := newTsvWriter()
	cols, err := rows.Columns()
	handleError(err)
	log.Debug("columns: %v", cols)

	values := make([]string, len(cols))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	if includeColumns == true {
		w.Write(cols)
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		handleError(err)
		if err := w.Write(values); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}

		w.Flush()
		if err := w.Error(); err != nil {
			log.Fatal(err)
		}
	}

	handleError(rows.Err())
}

func outputJson(rows *sql.Rows) {
	cols, err := rows.Columns()
	handleError(err)

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		jsonObj, _ := gabs.Consume(m)
		fmt.Println(jsonObj.String())
	}
}

func query(c string, q string) *sql.Rows {
	conn, err := sql.Open("mssql", c)
	handleError(err)
	defer conn.Close()

	stmt, err := conn.Prepare(opts.Query)
	handleError(err)
	defer stmt.Close()
	rows, err := stmt.Query()
	handleError(err)

	return rows
}

func getConnection() string {
	c := os.Getenv("MSSQL_CONN")

	if c == "" {
		handleError(errors.New("Environment variable MSSQL_CONN must be passed to connect to the DB"))
	}

	return c
}

// How can I live without you?
func handleError(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
