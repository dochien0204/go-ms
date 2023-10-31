package db

import (
	"fmt"
	"log"
	"order_svc/pkg/config"
	"order_svc/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func ConnectDB(c config.Config) Handler {

	host := c.Host
	user := c.User
	password := c.Password
	dbname := c.DatabaseName
	port := c.Port

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect DB")
	}

	fmt.Println("Connect DB Succesfully")

	db.AutoMigrate(&models.Order{})

	return Handler{db}

}
