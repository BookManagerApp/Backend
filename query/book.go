package query

import (
	"errors"
	"log"

	"github.com/BookManagerApp/Backend/model"

	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) ([]model.Book, error) {
    var books []model.Book
    if err := db.Find(&books).Error; err != nil {
        log.Printf("Error getting books: %v", err)
        return nil, err
    }
    return books, nil
}

func GetBookByID(db *gorm.DB, id int) (*model.Book, error) {
    var book model.Book
    if err := db.Where("id_book = ?", id).First(&book).Error; err != nil {
        return nil, err
    }
    return &book, nil
}

func PostBook(db *gorm.DB, book model.Book) error { 
	if err := db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, id int, updatedBook model.Book) error { 
    result := db.Model(&model.Book{}).Where("id_book = ?", id).Updates(updatedBook)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("tidak ada buku yang diperbarui")
    }
    return nil
}

func DeleteBook(db *gorm.DB, id int) error { 
    if err := db.Where("id_book = ?", id).Delete(&model.Book{}).Error; err != nil {
        return err
    }
    return nil
}

func GetGenres(db *gorm.DB) ([]string, error) {
	var genres []string
	if err := db.Model(&model.Book{}).Distinct("genre").Pluck("genre", &genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}
