package middleware

import (
	"Microservices/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.GetToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		}

		jwtToken := util.DecodeOpaqueueTokenToJWTToken(token)
		ctx.Header("Authorization", jwtToken)
		ctx.Next()

	}
}
