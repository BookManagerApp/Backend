package routes

import (
	"github.com/BookManagerApp/Backend/controller"
	"github.com/BookManagerApp/Backend/utils" 
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

	// Rute yang memerlukan autentikasi
	protected := app.Group("/protected", utils.AuthRequired) 
	protected.Get("/books", controller.GetBooks)
	protected.Get("/book/getbyid/:id", controller.GetBookByID)
	protected.Post("/book/post", controller.PostBook)
	protected.Put("/book/update/:id", controller.UpdateBook)
	protected.Delete("/book/delete/:id", controller.DeleteBook)
	protected.Get("/endpoint", controller.SomeProtectedHandler)
}

