package order

import (
	"Microservices/pkg/config"
	"Microservices/pkg/middleware"
	"Microservices/pkg/order/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.POST("", middleware.JWTMiddleware(), svc.CreateOrder)

}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
