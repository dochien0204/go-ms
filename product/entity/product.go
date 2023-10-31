package entity

import "time"

type Product struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Price float64
	Total int

	CreatedAt time.Time
	UpdatedAt time.Time
}
