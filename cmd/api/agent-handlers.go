package api

import (
	"net/http"

	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// Binding from JSON
type CreateAgentJson struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required,agent-status"`
	UserID int    `json:"user_id" binding:"required"`
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
		UserID: int32(createAgentJson.UserID),
	}

	agent, err := app.store.CreateAgent(ctx, newAgentArgs)

	if err != nil {
		if pgErr, ok := (err).(*pq.Error); ok {
			switch pgErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorJSON(err.Error()))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorJSON(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, agent)
}

type GetAgentQueryParams struct {
	ID int `uri:"id" binding:"required,min=1"`
}

func (app *Config) getAgent(ctx *gin.Context) {
	var agentId GetAgentQueryParams

	if err := ctx.ShouldBindUri(&agentId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON(err.Error()))
		return
	}

	agent, err := app.store.GetAgent(ctx, int64(agentId.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorJSON(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, agent)
}

type GetAgentsArgs struct {
	Page     int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=20"`
}

func (app *Config) getAgents(ctx *gin.Context) {
	var getAgentsArgs GetAgentsArgs

	if err := ctx.ShouldBindQuery(&getAgentsArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON("Provide the query params"))
		return
	}

	agents, err := app.store.ListAgents(ctx, db.ListAgentsParams{
		Limit:  getAgentsArgs.PageSize,
		Offset: getAgentsArgs.Page,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, agents)
}

type UpdateAgentArgs struct {
	Status string `json:"status" binding:"required,agent-status"`
}

func (app *Config) updateAgent(ctx *gin.Context) {
	var args UpdateAgentArgs
	var uriArgs GetAgentQueryParams

	if err := ctx.ShouldBindJSON(&args); err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON("Provide the status"))
		return
	}

	if err := ctx.ShouldBindUri(&uriArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON("Provide the agent id"))
		return
	}

	agent, err := app.store.UpdateAgent(ctx, db.UpdateAgentParams{
		Status: args.Status,
		ID:     int64(uriArgs.ID),
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, agent)

}

func (app *Config) deleteAgent(ctx *gin.Context) {
	var agent GetAgentQueryParams

	if err := ctx.ShouldBindUri(&agent); err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON("Provide the agent id"))
		return
	}
	err := app.store.DeleteAgent(ctx, int64(agent.ID))

	if err != nil {
		ctx.JSON(http.StatusConflict, errorJSON(err.Error()))
		return
	}
	response := map[string]string{
		"message": "Agent successfully deleted",
	}
	ctx.JSON(http.StatusOK, response)
}

func errorJSON(message string) gin.H {
	return gin.H{"message": message}
}
