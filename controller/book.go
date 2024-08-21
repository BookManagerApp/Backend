package controller

import (
	"net/http"

	"github.com/BookManagerApp/Backend/model"
	"github.com/BookManagerApp/Backend/query"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetBooks(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	books, err := query.GetBooks(db)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(books) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"code": http.StatusNotFound, "success": false, "status": "error", "message": "Data buku tidak ditemukan", "data": nil})
	}

	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    books,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID buku tidak ditemukan"})
	}

	db := c.Locals("db").(*gorm.DB)

	book, err := query.GetBookByID(db, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "data": book})
}

func PostBook(c *fiber.Ctx) error {
	var book model.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
	}

	db := c.Locals("db").(*gorm.DB)

	if err := query.PostBook(db, book); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan buku"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"code": http.StatusCreated, "success": true, "status": "success", "message": "Buku berhasil disimpan", "data": book})
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID buku tidak ditemukan"})
	}

	var updatedBook model.Book

	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
	}

	db := c.Locals("db").(*gorm.DB)

	if err := query.UpdateBook(db, id, updatedBook); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui buku"})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Buku berhasil diperbarui"})
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID buku tidak ditemukan"})
	}

	db := c.Locals("db").(*gorm.DB)

	if err := query.DeleteBook(db, id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus buku"})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Buku berhasil dihapus", "deleted_id": id})
}
