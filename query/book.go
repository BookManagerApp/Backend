package query

import (
	"github.com/BookManagerApp/Backend/model"
	"github.com/gocroot/model"
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

func UpdateBuku(db *gorm.DB, id string, updatedBuku model.Buku) error { // Memperbarui data buku dalam database berdasarkan ID
	if err := db.Model(&model.Buku{}).Where("id_buku = ?", id).Updates(updatedBuku).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBuku(db *gorm.DB, id string) error { // Menghapus data buku dari database berdasarkan ID
	if err := db.Delete(&model.Buku{}, id).Error; err != nil {
		return err
	}
	return nil
}
