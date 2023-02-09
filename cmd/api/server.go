package api

import (
	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/ShadrackAdwera/ticket-assignment/token"
	"github.com/ShadrackAdwera/ticket-assignment/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	store      db.TxStore
	router     *gin.Engine
	config     utils.Config
	tokenMaker token.TokenMaker
}

func NewServer(store db.TxStore, config utils.Config) (*Config, error) {
	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)

	if err != nil {
		return nil, err
	}

	router := gin.Default()
	server := Config{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
		router:     router,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("agent-status", isValidAgentStatus)
	}

	authenticatedRoutes := router.Group("/").Use(authMiddleware(tokenMaker))

	// routes
	router.POST("/auth/sign-up", server.signUp)
	authenticatedRoutes.POST("/agents", server.createAgent)
	authenticatedRoutes.GET("/agents/:id", server.getAgent)
	authenticatedRoutes.GET("/agents", server.getAgents)
	authenticatedRoutes.PATCH("/agents/:id", server.updateAgent)
	authenticatedRoutes.DELETE("/agents/:id", server.deleteAgent)

	server.router = router
	return &server, nil
}

func (app *Config) StartServer(port string) error {
	return app.router.Run(port)
}
