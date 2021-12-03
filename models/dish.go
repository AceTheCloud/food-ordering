package models

import "time"

type Dish struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	Created time.Time

	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`

	RestaurantRefer int        `json:"restaurant_id"`
	Restaurant      Restaurant `gorm:"foreignKey:RestaurantRefer"`
}
