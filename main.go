package main

import (
	"database/sql"
	"log"

	api "github.com/ShadrackAdwera/ticket-assignment/cmd/api"
	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/ShadrackAdwera/ticket-assignment/utils"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	dbInstance, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	// run migration
	runDbMiration(config.MigrationUrl, config.DbSource)

	store := db.NewTxStore(dbInstance)
	server, err := api.NewServer(store, config)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	err = server.StartServer(config.ServerAddress)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func runDbMiration(migrationUrl string, dbSource string) {
	m, err := migrate.New(migrationUrl, dbSource)
	if err != nil {
		log.Fatalf("failed to create a new migrate instance: %v", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrate up %v", err)
	}

	log.Println("Migration ran successfully . . . ")
}
