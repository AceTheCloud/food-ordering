package models

import "time"

type Dish struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	Created time.Time

	Name      string `json:"name"`
	Available bool   `gorm:"default:true"`
	Deleted   bool   `gorm:"default:false"`

	RestaurantRefer uint       `json:"restaurant_id"`
	Restaurant      Restaurant `gorm:"foreignKey:RestaurantRefer"`

	CuisineRefer uint    `json:"cuisine_id"`
	Cuisine      Cuisine `gorm:"foreignKey:CuisineRefer"`
}
