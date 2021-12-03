package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Email string
}
