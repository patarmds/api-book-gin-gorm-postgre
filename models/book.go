package models

import (
	"time"
)

type Book struct{
	ID	uint `json:"id" gorm:"primaryKey"`
	NameBook string `json:"name_book" gorm:"not null;type:varchar(255)"`
	Author string `json:"author" gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}