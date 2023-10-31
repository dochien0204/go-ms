package handler

import (
	"context"
	"fmt"
	"net/http"
	"product_svc/api/pb"
	"product_svc/infrastructure/define"
	"product_svc/usecase"
)

type Handler struct {
	service *usecase.Service
}

func NewHandler(service *usecase.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) FindOne(ctx context.Context, payload *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	product, err := h.service.FindProductById(int(payload.Id))
	if err != nil {
		return &pb.FindOneResponse{
			Status:  fmt.Sprint(http.StatusInternalServerError),
			Message: fmt.Sprintf(`%v`, define.INTERNAL_SERVER_ERROR),
		}, err
	}

	return &pb.FindOneResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: fmt.Sprintf(`%v`, define.SUCCESS),
		Data:    convertProductToPresenter(product),
	}, nil
}

func (h Handler) CreateProduct(ctx context.Context, payload *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return nil, nil
}
