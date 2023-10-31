package models

type Order struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	Quantity  int64 `json:"quantity"`
	ProductId int64 `json:"productId"`
	UserId    int64 `json:"userId"`
}
