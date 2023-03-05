package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/models"
	"github.com/soramon0/natrous/pkg/utils"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := utils.ReadJsonFile("users", &users); err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: users, Count: len(users)})
}
