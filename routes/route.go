package routes

import (
	"github.com/BookManagerApp/Backend/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(app *fiber.App) {
	app.Get("/books", controller.GetBooks)
	app.Get("/book/getbyid/:id", controller.GetBookByID)
	app.Post("/book/post", controller.PostBook)
	app.Put("/book/update/:id", controller.UpdateBook)
	app.Delete("/book/delete/:id", controller.DeleteBook)
}

