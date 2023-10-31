package repository

import (
	"fmt"
	"product_svc/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "postgres"
	PASSWORD = "root"
	DBNAME   = "product_svc"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", HOST, USER, PASSWORD, DBNAME, PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.Product{})

	return db, nil
}
