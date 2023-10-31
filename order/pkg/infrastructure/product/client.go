package client

import (
	"context"
	"fmt"
	pb "order_svc/pkg/pb/product"

	"google.golang.org/grpc"
)

type ProductServicesClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServicesClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ProductServicesClient{
		Client: pb.NewProductServiceClient(cc),
	}

	return c
}

func (c *ProductServicesClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return c.Client.FindOne(context.Background(), req)
}
