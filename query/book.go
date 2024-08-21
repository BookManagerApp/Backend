package query

import (
	"github.com/BookManagerApp/Backend/model"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) ([]model.Book, error) { 
	var book []model.Book
	if err := db.Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetBookByID(db *gorm.DB, id string) (model.Book, error) {
	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		return book, err
	}
	return book, nil
}

func PostBook(db *gorm.DB, book model.Book) error { 
	if err := db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, id string, updatedBook model.Book) error {
	if err := db.Model(&model.Book{}).Where("id = ?", id).Updates(updatedBook).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(db *gorm.DB, id string) error { 
	if err := db.Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
