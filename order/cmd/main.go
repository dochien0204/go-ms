package main

import (
	"fmt"
	"log"
	"net"
	"order_svc/pkg/config"
	"order_svc/pkg/db"
	client "order_svc/pkg/infrastructure/product"
	pb "order_svc/pkg/pb/order"
	"order_svc/pkg/services"

	"google.golang.org/grpc"
)

const (
	PRODUCT_SERVICE_PORT = "localhost:8001"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config:", err)
	}

	h := db.ConnectDB(c)

	lis, err := net.Listen("tcp", ":8002")

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order svc on: 8002")

	productService := client.InitProductServiceClient(PRODUCT_SERVICE_PORT)
	s := services.Server{
		H:               h,
		ProductServices: productService,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
