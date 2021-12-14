package models

import "time"

type Cuisine struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name    string
	Deleted bool `gorm:"default:false"`

	Created time.Time
}
