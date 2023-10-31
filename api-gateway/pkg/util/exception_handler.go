package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BasicResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HandlerException(ctx *gin.Context, statusCode int, err error) {
	errorMessage := BasicResponse{
		Status:  fmt.Sprint(statusCode),
		Message: err.Error(),
	}

	ctx.AbortWithStatusJSON(statusCode, errorMessage)
}
