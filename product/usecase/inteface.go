package usecase

import (
	"product_svc/entity"
)

type ProductRepository interface {
	FindById(id int) (*entity.Product, error)
}

type UseCase interface {
	FindProductById(id int) (*entity.Product, error)
}
