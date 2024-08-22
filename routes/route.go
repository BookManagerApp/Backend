package routes

import (
	"github.com/BookManagerApp/Backend/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Rute untuk buku
	app.Get("/books", controller.GetBooks)
	app.Get("/book/getbyid/:id", controller.GetBookByID)
	app.Post("/book/post", controller.PostBook)
	app.Put("/book/update/:id", controller.UpdateBook)
	app.Delete("/book/delete/:id", controller.DeleteBook)
	app.Get("/genres", controller.GetGenres)

	// Rute untuk pengguna
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
}
