package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/models"
	"github.com/soramon0/natrous/pkg/repository"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := repository.GetUsers()
	if err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: users, Count: len(*users)})
}

func GetUser(c *fiber.Ctx) error {
	user, err := repository.GetUserByID(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: user})
}

func CreateUser(c *fiber.Ctx) error {
	user, err := repository.CreateUser()
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: user})
}
