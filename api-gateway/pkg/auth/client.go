package auth

import (
	"Microservices/pkg/auth/pb"
	"Microservices/pkg/config"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServicesClient
}

func InitServiceClient(c *config.Config) pb.AuthServicesClient {

	client, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	fmt.Println("URL:", c.AuthSvcUrl)
	if err != nil {
		fmt.Println("Cannot connect:", err)
	}

	return pb.NewAuthServicesClient(client)
}
