package models

import "time"

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time

	DishRefer int  `json:"product_id"`
	Dish      Dish `gorm:"foreignKey:DishRefer"`

	UserRefer int  `json:"user_id"`
	User      User `gorm:"foreignKey:UserRefer"`
}
