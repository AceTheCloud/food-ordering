package models

import "time"

type Restaurant struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	Created time.Time

	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}
