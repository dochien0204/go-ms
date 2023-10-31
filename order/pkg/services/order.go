package services

import (
	"context"
	"net/http"
	"order_svc/pkg/db"
	client "order_svc/pkg/infrastructure/product"
	"order_svc/pkg/models"
	pb "order_svc/pkg/pb/order"
	"order_svc/pkg/util"
)

type Server struct {
	H               db.Handler
	ProductServices client.ProductServicesClient
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	product, err := s.ProductServices.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	//Parse token
	claims, err := util.ParseAccessToken(req.JwtToken)
	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		}, nil
	}

	order := models.Order{
		Quantity:  req.Quantity,
		ProductId: product.GetData().Id,
		UserId:    claims.UserId,
	}

	s.H.DB.Create(&order)

	return &pb.CreateOrderResponse{
		Status: http.StatusOK,
		Id:     order.Id,
	}, nil
}
