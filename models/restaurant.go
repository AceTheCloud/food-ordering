package models

import "time"

type Restaurant struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name string
	City string

	OpeningTime time.Time `json:"opening_time" gorm:"opening_time"`
	ClosingTime time.Time `json:"closing_time" gorm:"closing_time"`

	Open    bool `gorm:"default:true"`
	Deleted bool `gorm:"default:false"`

	CreatedAt time.Time
}
