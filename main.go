package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/BookManagerApp/Backend/config"
	"github.com/BookManagerApp/Backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func main() {
    app := fiber.New()

    // Membuat koneksi ke database
    db := config.CreateDBConnection()

    // Jalankan SQL file untuk menginisialisasi database
    importSQLFile("./book.sql", db) // Pastikan path sesuai dengan lokasi file di backend

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
    routes.SetupBookRoutes(app)

    // Menjalankan server Fiber pada port 3000
    if err := app.Listen(":3000"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

// Fungsi untuk mengimpor file SQL ke database menggunakan gorm.DB
func importSQLFile(filepath string, db *gorm.DB) {
    // Membaca file SQL
    sqlFile, err := ioutil.ReadFile(filepath)
    if err != nil {
        log.Fatalf("Gagal membaca file SQL: %v", err)
    }

    // Pisahkan perintah SQL yang dipisahkan oleh ";"
    queries := strings.Split(string(sqlFile), ";")

    // Eksekusi setiap query secara berurutan
    for _, query := range queries {
        trimmedQuery := strings.TrimSpace(query)
        if trimmedQuery != "" {
            err := db.Exec(trimmedQuery).Error
            if err != nil {
                log.Printf("Gagal menjalankan query: %s, Error: %v", trimmedQuery, err)
            } else {
                fmt.Printf("Query berhasil dijalankan: %s\n", trimmedQuery)
            }
        }
    }

    log.Println("SQL script berhasil dijalankan.")
}
