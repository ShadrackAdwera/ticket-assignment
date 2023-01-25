package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple-bank?sslmode=disable" // TODO: Load via env
)

var testQuery *Queries

func TestMain(m *testing.M) {
	sqlDb, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	testQuery = New(sqlDb)
	os.Exit(m.Run())
}
