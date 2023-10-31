package order

import (
	"Microservices/pkg/config"
	"Microservices/pkg/order/pb"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Cannot connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
