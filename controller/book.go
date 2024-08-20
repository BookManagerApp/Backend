package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBooks(respw http.ResponseWriter, req *http.Request) {
	books, err := atdb.GetAllDoc[[]model.Book](config.Mongoconn, "books", bson.M{})
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, books)
}

func GetOneBook(respw http.ResponseWriter, req *http.Request) {
	// Ambil ID dari URL
	idParam := req.URL.Query().Get("id")
	if idParam == "" {
		helper.WriteJSON(respw, http.StatusBadRequest, "Missing book ID")
		return
	}

	// Konversi ID dari string ke ObjectID
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, "Invalid book ID")
		return
	}

	// Buat filter untuk mencari dokumen dengan ID
	filter := bson.M{"_id": id}

	// Ambil satu dokumen buku
	book, err := atdb.GetOneDoc[model.Book](config.Mongoconn, "books", filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helper.WriteJSON(respw, http.StatusNotFound, "Book not found")
		} else {
			helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Kembalikan dokumen buku dalam format JSON
	helper.WriteJSON(respw, http.StatusOK, book)
}

func PostBook(respw http.ResponseWriter, req *http.Request) {
	var newBook model.Book
	if err := json.NewDecoder(req.Body).Decode(&newBook); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}
	newBook.ID = primitive.NewObjectID()

	// Validasi data buku
	if newBook.Title == "" || newBook.Author == "" || newBook.Genre == "" || newBook.PublishedYear == 0 {
		helper.WriteJSON(respw, http.StatusBadRequest, "All fields (Title, Author, PublishedYear, Genre) are required")
		return
	}

	if _, err := atdb.InsertOneDoc(config.Mongoconn, "books", newBook); err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(respw, http.StatusOK, newBook)
}

func UpdateBook(respw http.ResponseWriter, req *http.Request) {
    var book model.Book
    err := json.NewDecoder(req.Body).Decode(&book)
    if err != nil {
        helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
        return
    }

    // Validasi bahwa ID buku tidak boleh kosong
    if book.ID == primitive.NilObjectID {
        helper.WriteJSON(respw, http.StatusBadRequest, "Book ID is required")
        return
    }

    // Definisikan filter untuk menemukan buku berdasarkan ID
    filter := bson.M{"_id": book.ID}

    // Definisikan update dengan set data baru
    update := bson.M{
        "$set": bson.M{
            "title":        book.Title,
            "author":       book.Author,
            "publishedYear": book.PublishedYear,
            "genre":        book.Genre,
        },
    }

    // Update buku di MongoDB dan tangkap kedua nilai yang dikembalikan
    result, err := atdb.UpdateOneDoc(config.Mongoconn, "books", filter, update)
    if err != nil {
        helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
        return
    }

    // Cek hasil update untuk memastikan ada dokumen yang diubah
    if result.MatchedCount == 0 {
        helper.WriteJSON(respw, http.StatusNotFound, "Book not found")
        return
    }

    // Kirim respons sukses
    helper.WriteJSON(respw, http.StatusOK, "Book successfully updated")
}


func DeleteBook(respw http.ResponseWriter, req *http.Request) {
	var book model.Book
	if err := json.NewDecoder(req.Body).Decode(&book); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, err.Error())
		return
	}

	// Validasi bahwa ID buku tidak boleh kosong
	if book.ID == primitive.NilObjectID {
		helper.WriteJSON(respw, http.StatusBadRequest, "Book ID is required")
		return
	}

	// Definisikan filter untuk menemukan buku berdasarkan ID
	filter := bson.M{"_id": book.ID}

	// Hapus buku dari MongoDB dan tangkap kedua nilai yang dikembalikan
	result, err := atdb.DeleteOneDoc(config.Mongoconn, "books", filter)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, err.Error())
		return
	}

	// Cek hasil delete untuk memastikan ada dokumen yang dihapus
	if result.DeletedCount == 0 {
		helper.WriteJSON(respw, http.StatusNotFound, "Book not found")
		return
	}

	// Kirim respons sukses
	helper.WriteJSON(respw, http.StatusOK, "Book successfully deleted")
}

