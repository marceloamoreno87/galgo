package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name" gorm:"index"`
	Username  string    `json:"username" gorm:"unique" gorm:"index"`
	Email     string    `json:"email" gorm:"unique" gorm:"index"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at" gorm:"index"`
}
