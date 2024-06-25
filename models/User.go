package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstName" form:"firstName"`
	LastName  string    `json:"lastName" form:"lastName"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at"`
}
