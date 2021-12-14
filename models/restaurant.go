package models

import "time"

type Restaurant struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name string
	City string

	OpeningTime uint8 `json:"opening_time" gorm:"opening_time"`
	ClosingTime uint8 `json:"closing_time" gorm:"closing_time"`

	Open    bool `gorm:"default:true"`
	Deleted bool `gorm:"default:false"`

	CreatedAt time.Time
}
