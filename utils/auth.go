package utils

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// AuthRequired memeriksa autentikasi JWT
func AuthRequired(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	}

	role := claims["role"].(string)
	c.Locals("role", role)

	return c.Next()
}
