package model

import "time"

type Book struct {
    ID            int    `gorm:"primaryKey;column:id_book" json:"id_book"`
    Title         string `gorm:"column:title" json:"title"`
    Author        string `gorm:"column:author" json:"author"`
    PublishedYear int    `gorm:"column:publishedyear" json:"publishedyear"`
    Genre         string `gorm:"column:genre" json:"genre"`
}

type Users struct {
	IDUser    uint      `gorm:"primaryKey;column:id_user" json:"id_user"`
	Password  string    `gorm:"column:password" json:"password"`
	Email     string    `gorm:"column:email" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}