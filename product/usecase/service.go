package usecase

import (
	"product_svc/entity"
)

type Service struct {
	productRepo ProductRepository
}

func NewService(productRepo ProductRepository) *Service {
	return &Service{
		productRepo: productRepo,
	}
}

func (s Service) FindProductById(id int) (*entity.Product, error) {
	return s.productRepo.FindById(id)
}
