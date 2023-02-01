package api

import (
	"net/http"

	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type CreateAgentJson struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (app *Config) createAgent(ctx *gin.Context) {
	var createAgentJson CreateAgentJson

	if err := ctx.ShouldBindJSON(&createAgentJson); err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON(err.Error()))
		return
	}
	newAgentArgs := db.CreateAgentParams{
		Name:   createAgentJson.Name,
		Status: createAgentJson.Status,
	}

	agent, err := app.store.CreateAgent(ctx, newAgentArgs)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorJSON(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, agent)
}

func errorJSON(message string) gin.H {
	return gin.H{"message": message}
}
