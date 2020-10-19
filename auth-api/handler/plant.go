package handler

import (
	"fmt"
	"golang_fiber_auth/auth-api/model"

	"github.com/gofiber/fiber/v2"
)

// GetPlantsHandler used to get the plant records
func (auth *Auth) GetPlantsHandler(c *fiber.Ctx) error {

	user := model.User{}
	err := auth.db.Where("id = ?", c.Locals("user_id")).Find(&user).Error
	if err != nil {
		auth.logger.Printf("[ERROR] User record not found. Reason: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("User record not found. Reason: %v", err),
		})
	}

	plants := []*model.Plant{
		&model.Plant{
			Name:        "Mango",
			Category:    "Fruit",
			Price:       300.00,
			Description: "Sweetest Fruit",
			Avatar:      "mango.png",
			User:        user,
		},
	}

	err = auth.db.Create(&plants).Error
	if err != nil {
		auth.logger.Printf("[ERROR] Plant record creation failed. Reason: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Plant record creation failed. Reason: %v", err),
		})
	}
	auth.db.Preload("User").Find(&plants)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"plants": plants,
	})

}
