package utils

import (
    "github.com/gofiber/fiber/v2"
    "strings"
    "log"
)

// AuthRequired memeriksa autentikasi JWT
func AuthRequired(c *fiber.Ctx) error {
    token := c.Get("Authorization")
    if token == "" {
        return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
    }

    token = strings.TrimPrefix(token, "Bearer ")
    log.Printf("Received token: %s", token) // Tambahkan log

    _, claims, err := ParseToken(token)
    if err != nil {
        log.Printf("Token parsing error: %v", err) // Tambahkan log
        return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
    }

    c.Locals("user", claims)
    return c.Next()
}
