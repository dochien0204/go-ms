package main

import (
	"fmt"
	"log"
	"net"
	"order_svc/pkg/config"
	"order_svc/pkg/db"
	"order_svc/pkg/pb"
	"order_svc/pkg/services"

	"google.golang.org/grpc"
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

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
