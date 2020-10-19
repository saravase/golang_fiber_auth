package handler

import (
	"fmt"
	"golang_fiber_auth/auth-api/model"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Auth struct {
	logger *log.Logger
	db     *gorm.DB
}

// Initialize Auth struct properties
func NewAuth(logger *log.Logger, db *gorm.DB) *Auth {
	return &Auth{
		logger: logger,
		db:     db,
	}
}

// SignupHandler used to handle the signup request
func (auth *Auth) SignupHandler(c *fiber.Ctx) error {

	user := &model.User{}

	if err := c.BodyParser(user); err != nil {
		auth.logger.Printf("[ERROR] JSON parsing failed. Reason: %v", err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": fmt.Sprintf("JSON parsing failed. Reason: %v", err),
		})
	}

	password, err := EncryptPassword(user.Password)
	if err != nil {
		auth.logger.Printf("[ERROR] Password hashing failed. Reason: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Password hashing failed. Reason: %v", err),
		})
	}

	user.Password = password

	err = auth.db.Create(user).Error
	if err != nil {
		auth.logger.Printf("[ERROR] User creation failed. Reason: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("User creation failed. Reason: %v", err),
		})
	}

	auth.db.Where("email = ?", user.Email).First(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": user.ID,
	})
}

// SigninHandler used to handle the signin request
func (auth *Auth) SigninHandler(c *fiber.Ctx) error {

	loginUser := &model.User{}

	if err := c.BodyParser(loginUser); err != nil {
		auth.logger.Printf("[ERROR] JSON parsing failed. Reason: %v", err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": fmt.Sprintf("JSON parsing failed. Reason: %v", err),
		})
	}

	user := &model.User{}
	result := auth.db.Where("email = ?", loginUser.Email).First(user)
	if result.RowsAffected > 0 {
		if !ValidatePassword(loginUser.Password, user.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User credentials wrong",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not exist",
		})
	}

	err := CreateToken(c, user.ID)
	if err != nil {
		auth.logger.Printf("[ERROR] Token creation failed. Reason: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Token creation failed. Reason: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"name":  user.Name,
		"email": user.Email,
	})
}

// SignoutHandler used to handle the signout request
func (auth *Auth) SignoutHandler(c *fiber.Ctx) error {

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Expires = time.Now()
	c.Cookie(cookie)
	return c.SendStatus(fiber.StatusOK)

}
