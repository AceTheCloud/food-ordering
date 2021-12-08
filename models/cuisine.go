package models

import "time"

type Cuisine struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	Created time.Time

	Name string
}
