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
    // Cek hasil
    for _, book := range books {
        log.Printf("Book ID: %d, Title: %s", book.ID, book.Title)
    }
    return books, nil
}


func GetBookByID(db *gorm.DB, id int) (*model.Book, error) { // Ubah ke int
    var book model.Book
    if err := db.First(&book, id).Error; err != nil {
        return nil, err
    }
    return &book, nil
}

func PostBook(db *gorm.DB, book model.Book) error {
    if err := db.Create(&book).Error; err != nil {
        return err
    }
    // Pastikan ID sudah diisi
    log.Printf("Book ID after creation: %d", book.ID)
    return nil
}


func UpdateBook(db *gorm.DB, id int, updatedBook model.Book) error {
    result := db.Model(&model.Book{}).Where("id = ?", id).Updates(updatedBook)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("tidak ada buku yang diperbarui")
    }
    // Cek hasil update
    log.Printf("Rows affected: %d", result.RowsAffected)
    return nil
}

func DeleteBook(db *gorm.DB, id int) error {
    result := db.Delete(&model.Book{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("tidak ada buku yang dihapus")
    }
    // Cek hasil delete
    log.Printf("Rows affected: %d", result.RowsAffected)
    return nil
}

