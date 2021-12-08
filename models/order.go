package models

import "time"

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time

	DishRefer uint `json:"dish_id"`
	Dish      Dish `gorm:"foreignKey:DishRefer"`

	UserRefer uint `json:"user_id"`
	User      User `gorm:"foreignKey:UserRefer"`

	Status string
}
