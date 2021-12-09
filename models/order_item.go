package models

import "time"

type OrderItem struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderId string `json:"order_id"`

	DishId uint `json:"dish_id"`

	Quantity uint `json:"quantity"`

	CreatedAt time.Time
}
