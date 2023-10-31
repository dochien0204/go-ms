package handler

import (
	"Microservices/pkg/config"
	"Microservices/pkg/product/pb"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Cannot connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
