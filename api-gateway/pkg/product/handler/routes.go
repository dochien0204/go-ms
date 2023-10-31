package handler

import (
	"Microservices/pkg/config"
	"Microservices/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.GET("", middleware.JWTMiddleware(), svc.FindOne)

}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	FindProduct(ctx, svc.Client)
}
