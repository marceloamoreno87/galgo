package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"index"`
	Content   string    `json:"content" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at" gorm:"index"`
}
