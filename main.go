package main

import (
	"log"
	"os"

	"github.com/BookManagerApp/Backend/config"
	"github.com/BookManagerApp/Backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()

    // Membuat koneksi ke database
    db := config.CreateDBConnection()

    // Mengatur middleware logger
    app.Use(logger.New(logger.Config{
        Format: "${status} - ${method} ${path}\n",
    }))

    // Mengatur middleware CORS
    app.Use(cors.New(cors.Config{
        AllowHeaders: "*",
        AllowOrigins: "*",
        AllowMethods: "GET, POST, PUT, DELETE",
    }))

    // Menyimpan koneksi database ke context Fiber
    app.Use(func(c *fiber.Ctx) error {
        c.Locals("db", db)
        return c.Next()
    })

    // Menyiapkan rute buku
    routes.SetupRoutes(app)

    // Menambahkan rute untuk menguji apakah server berjalan
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running.")
	})

	// Mengambil port dari variabel lingkungan
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port jika tidak ada variabel PORT
	}

	// Menjalankan server Fiber pada port yang ditentukan
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
