package model

import "time"

// Book представляет модель книги в библиотеке
type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Author      string    `json:"author" gorm:"not null"`
	ISBN        string    `json:"isbn" gorm:"unique"`
	Description string    `json:"description"`
	Year        int       `json:"year"`
	Publisher   string    `json:"publisher"`
	Available   bool      `json:"available" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BookCreate представляет структуру для создания новой книги
type BookCreate struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	Description string `json:"description"`
	Year        int    `json:"year" binding:"required"`
	Publisher   string `json:"publisher"`
} 