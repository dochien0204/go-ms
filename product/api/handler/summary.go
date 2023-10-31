package handler

import (
	"product_svc/api/pb"
	"product_svc/entity"
)

func convertProductToPresenter(data *entity.Product) *pb.FindOneData {
	return &pb.FindOneData{
		Id:    int64(data.Id),
		Name:  data.Name,
		Price: int64(data.Price),
		Total: int64(data.Total),
	}
}
