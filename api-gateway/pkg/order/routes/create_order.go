package routes

import (
	"Microservices/pkg/order/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	ProductId     int64  `json:"productId"`
	Quantity      int64  `json:"quantity"`
	OpaqueueToken string `json:"jwtToken"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		JwtToken:  body.OpaqueueToken,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
