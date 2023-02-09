package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ShadrackAdwera/ticket-assignment/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "authorization"
	authTypeBearer         = "bearer"
	authPayloadKey         = "authPayload"
)

func authMiddleware(tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("invalid authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorJSON(err.Error()))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorJSON(err.Error()))
			return
		}
		authType := strings.ToLower(fields[0])
		if authType != authTypeBearer {
			err := errors.New(" authorization type is not supported")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorJSON(err.Error()))
			return
		}
		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorJSON(err.Error()))
			return
		}

		ctx.Set(authPayloadKey, payload)
		ctx.Next()
	}
}
