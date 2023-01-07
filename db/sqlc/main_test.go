package db

import (
	"database/sql"
	"jnwanya/simplebank/db/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:p@ssw0rd@localhost:6543/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatalln("Cannot load app configs.", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalln("Cannot connect to the DB.", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
