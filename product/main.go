package main

import (
	"fmt"
	"log"
	"net"
	"product_svc/api/handler"
	"product_svc/api/pb"
	"product_svc/infrastructure/define"
	"product_svc/infrastructure/repository"
	"product_svc/usecase"

	"google.golang.org/grpc"
)

func main() {

	db, err := repository.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed at config: ", err)
	}

	lis, err := net.Listen("tcp", define.PORT_APPLICATION)

	if err != nil {
		log.Fatalln("Failed To listing:", err)
	}

	fmt.Println("Product Svc on:", define.PORT_APPLICATION)

	//repo
	productRepo := repository.NewProductRepository(db)

	//service
	productService := usecase.NewService(productRepo)

	//Handler
	productHandler := handler.NewHandler(productService)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
