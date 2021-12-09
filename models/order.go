package models

import (
	"time"

	"github.com/acethecloud/food-ordering/enums"
)

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time

	UserRefer uint `json:"user_id"`
	OrderedBy User `gorm:"foreignKey:UserRefer"`

	TotalPrice uint `json:"total_price"`

	Status enums.Status `json:"status" gorm:"default:1"`
}
