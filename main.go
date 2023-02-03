package main

import (
	"database/sql"
	"log"

	api "github.com/ShadrackAdwera/ticket-assignment/cmd/api"
	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	// config, err := utils.LoadConfig(".")

	// if err != nil {
	// 	log.Fatalf(err.Error())
	// 	return
	// }

	dbInstance, err := sql.Open("postgresql://root:password@localhost:5432/ticket-assignment?sslmode=disable", "postgres")
	//dbInstance, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	store := db.NewTxStore(dbInstance)
	server := api.NewServer(store)

	err = server.StartServer("0.0.0.0:5000")
	//err = server.StartServer(config.ServerAddress)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}
