package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/ShadrackAdwera/ticket-assignment/utils"
	_ "github.com/lib/pq"
)

var testQuery *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../..")

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	testDb, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	testQuery = New(testDb)
	os.Exit(m.Run())
}
