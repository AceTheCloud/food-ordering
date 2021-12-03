package models

import "time"

type Restaurant struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time

	Name     string
	City     string
	Cuisines []string `gorm:"type:string[]"` //explicitly tell gorm about array type
}
