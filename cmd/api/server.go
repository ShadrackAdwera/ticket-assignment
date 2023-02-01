package api

import (
	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Config struct {
	store  *db.TxStore
	router *gin.Engine
}

func NewServer(store *db.TxStore) *Config {
	server := Config{
		store: store,
	}
	router := gin.Default()

	// routes
	router.POST("/agents", server.createAgent)

	server.router = router
	return &server
}

func (app *Config) StartServer(port string) error {
	return app.router.Run(port)
}
