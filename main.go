package main

import (
	"database/sql"
	"log"

	api "github.com/ShadrackAdwera/ticket-assignment/cmd/api"
	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/ShadrackAdwera/ticket-assignment/utils"

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

	store := db.NewTxStore(dbInstance)
	server, err := api.NewServer(store, config)

	err = server.StartServer(config.ServerAddress)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}
