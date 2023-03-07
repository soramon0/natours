package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/models"
)

type Users struct {
	service models.UserService
	log     *log.Logger
}

// New Users is used to create a new Users controller.
func NewUsers(us models.UserService, l *log.Logger) *Users {
	return &Users{
		service: us,
		log:     l,
	}
}

func (u *Users) GetUsers(c *fiber.Ctx) error {
	users, err := u.service.Find()
	if err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: users, Count: len(*users)})
}

func (u *Users) GetUser(c *fiber.Ctx) error {
	user, err := u.service.ByID(c.Params("id"))
	if err != nil {
		u.log.Println(err)
		return &fiber.Error{Code: fiber.StatusNotFound, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: user})
}

func (u *Users) CreateUser(c *fiber.Ctx) error {
	user, err := u.service.Create()
	if err != nil {
		u.log.Println(err)
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: user})
}
