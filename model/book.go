package model

type Book struct {
	ID            int    `gorm:"primaryKey;column:id" json:"id"` 
	Title         string `gorm:"column:title" json:"title"`
	Author        string `gorm:"column:author" json:"author"`
	PublishedYear int    `gorm:"column:publishedyear" json:"publishedyear"`
	Genre         string `gorm:"column:genre" json:"genre"`
}
