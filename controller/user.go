package controller

import (
	"github.com/BookManagerApp/Backend/model"
	"github.com/BookManagerApp/Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

// Register handler
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

	db := c.Locals("db").(*gorm.DB)
	if result := db.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Login handler
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

	// Here you should generate a JWT token and return it in the response
	return c.Status(fiber.StatusOK).SendString("Login successful")
}
