package services

import (
	"context"
	"fmt"
	"net/http"
	"order_svc/pkg/db"
	"order_svc/pkg/models"
	"order_svc/pkg/pb"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	token := req.Token
	fmt.Println(token)

	order := models.Order{
		Quantity:  req.Quantity,
		ProductId: req.ProductId,
	}

	s.H.DB.Create(&order)

	return &pb.CreateOrderResponse{
		Status: http.StatusOK,
		Id:     order.Id,
	}, nil
}
