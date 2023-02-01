package api

import (
	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Config struct {
	store  db.TxStore
	router *gin.Engine
}

func NewServer(store db.TxStore) *Config {
	server := Config{
		store: store,
	}
	router := gin.Default()

	// routes
	router.POST("/agents", server.createAgent)
	router.GET("/agents/:id", server.getAgent)
	router.GET("/agents", server.getAgents)
	router.PATCH("/agents/:id", server.updateAgent)
	router.DELETE("/agents/:id", server.deleteAgent)

	server.router = router
	return &server
}

func (app *Config) StartServer(port string) error {
	return app.router.Run(port)
}
