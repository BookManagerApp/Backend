package controller

import (
	"github.com/BookManagerApp/Backend/model"
	"github.com/BookManagerApp/Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

// Register
func Register(c *fiber.Ctx) error {
	var user model.Users
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to hash password")
	}

	user.Password = hashedPassword

	// Default ke "pengguna" jika tidak ditentukan
	if user.Role == "" {
		user.Role = "user"
	}

	db := c.Locals("db").(*gorm.DB)
	// Check if email already exists
	var existingUser model.Users
	if result := db.Where("email = ?", user.Email).First(&existingUser); result.Error == nil {
		return c.Status(fiber.StatusConflict).SendString("Email already registered")
	}

	if result := db.Create(&user); result.Error != nil {
		log.Printf("Failed to register user: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Login
func Login(c *fiber.Ctx) error {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	db := c.Locals("db").(*gorm.DB)
	var user model.Users
	if result := db.Where("email = ?", loginData.Email).First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid email or password")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to query user")
	}

	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid email or password")
	}

	// Generate JWT token with role
	token, err := utils.GenerateToken(user.IDUser, user.Email, user.Role)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate token")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"role":  user.Role,
	})
}

// SomeProtectedHandler, handler untuk rute yang dilindungi
func SomeProtectedHandler(c *fiber.Ctx) error {
    return c.SendString("This is a protected endpoint")
}

