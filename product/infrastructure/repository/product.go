package repository

import (
	"product_svc/entity"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r ProductRepository) FindById(id int) (*entity.Product, error) {
	product := &entity.Product{}
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepository) CreateProduct(data *entity.Product) error {
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}
