package handler

import (
	"Microservices/pkg/product/pb"
	"Microservices/pkg/util"
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	productId := ctx.Query("productId")
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, errors.New("Bad request"))
		return
	}

	_, err = util.GetToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(productIdInt),
	})

	if err != nil {
		util.HandlerException(ctx, http.StatusInternalServerError, errors.New("Internal server error"))
		return
	}

	ctx.JSON(http.StatusOK, res)

}
