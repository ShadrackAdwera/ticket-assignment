package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuery *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://adwera:mdcclxxvi@localhost:5432/ticket-assignment?sslmode=disable"
)

func TestMain(m *testing.M) {
	db, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	testQuery = New(db)
	os.Exit(m.Run())
}
