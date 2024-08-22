package controller

import (
	"net/http"
	"strconv"

	"github.com/BookManagerApp/Backend/model"
	"github.com/BookManagerApp/Backend/query"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetBooks(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var books []model.Book
	if err := db.Find(&books).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    books,
	})
}



func GetBookByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID buku tidak valid"})
	}

	db := c.Locals("db").(*gorm.DB)

	book, err := query.GetBookByID(db, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if book == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"success": false,
			"status":  "error",
			"message": "Buku tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    book,
	})
}

func PostBook(c *fiber.Ctx) error {
    var book model.Book

    if err := c.BodyParser(&book); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
    }

    db := c.Locals("db").(*gorm.DB)

    if err := query.PostBook(db, book); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save book"})
    }

    return c.Status(http.StatusCreated).JSON(fiber.Map{
        "code":    http.StatusCreated,
        "success": true,
        "status":  "success",
        "message": "Book saved successfully",
        "data":    book,  // Pastikan `book` sudah termasuk field ID
    })
}

func UpdateBook(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID buku tidak valid"})
    }

    var updatedBook model.Book

    if err := c.BodyParser(&updatedBook); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
    }

    db := c.Locals("db").(*gorm.DB)

    if err := query.UpdateBook(db, id, updatedBook); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui buku"})
    }

    return c.JSON(fiber.Map{
        "code":    http.StatusOK,
        "success": true,
        "status":  "success",
        "message": "Buku berhasil diperbarui",
        "data":    updatedBook, // Pastikan ID juga ikut
    })
}

func DeleteBook(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID buku tidak valid"})
    }

    db := c.Locals("db").(*gorm.DB)

    if err := query.DeleteBook(db, id); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus buku"})
    }

    return c.JSON(fiber.Map{
        "code":    http.StatusOK,
        "success": true,
        "status":  "success",
        "message": "Buku berhasil dihapus",
        "deleted_id": id, // ID yang dihapus
    })
}
