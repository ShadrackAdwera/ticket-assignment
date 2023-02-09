package api

import (
	"net/http"
	"time"

	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/ShadrackAdwera/ticket-assignment/utils"
	"github.com/gin-gonic/gin"
)

type CreateUserJSON struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserStruct struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
type UserResponse struct {
	User        *UserStruct `json:"user"`
	AccessToken string      `json:"access_token"`
}

func (app *Config) signUp(ctx *gin.Context) {
	var user CreateUserJSON

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON(err.Error()))
		return
	}

	password, err := utils.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorJSON(err.Error()))
		return
	}

	newUser, err := app.store.CreateUser(ctx, db.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorJSON(err.Error()))
		return
	}

	tkn, err := app.tokenMaker.CreateToken(newUser.Username, newUser.ID, app.config.TokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorJSON(err.Error()))
		return
	}

	authUser := &UserResponse{
		User: &UserStruct{
			ID:        newUser.ID,
			Username:  newUser.Username,
			Email:     newUser.Email,
			CreatedAt: newUser.CreatedAt,
		},
		AccessToken: tkn,
	}
	ctx.JSON(http.StatusCreated, authUser)
}
