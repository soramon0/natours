package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/models"
)

type Tours struct {
	ts models.TourService
	l  *log.Logger
}

// New Users is used to create a new Users controller.
func NewTours(ts models.TourService, l *log.Logger) *Tours {
	return &Tours{
		ts: ts,
		l:  l,
	}
}

func (t *Tours) GetTours(c *fiber.Ctx) error {
	tours, err := t.ts.Find()
	if err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: tours, Count: len(*tours)})
}

func (t *Tours) GetTour(c *fiber.Ctx) error {
	tour, err := t.ts.ByID(c.Params("id"))
	if err != nil {
		t.l.Println(err)
		return &fiber.Error{Code: fiber.StatusNotFound, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: tour})
}

func (t *Tours) CreateTour(c *fiber.Ctx) error {
	payload := struct {
		Name string `json:"name"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	if payload.Name == "" {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Name is required"}
	}

	tour := models.Tour{Name: payload.Name}
	return c.JSON(models.APIResponse{Data: tour})
}

func (t *Tours) UpdateTour(c *fiber.Ctx) error {
	tour, err := t.ts.ByID(c.Params("id"))
	if err != nil {
		t.l.Println(err)
		return &fiber.Error{Code: fiber.StatusNotFound, Message: err.Error()}
	}

	payload := models.Tour{}
	if err := c.BodyParser(&payload); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	return c.Status(fiber.StatusCreated).JSON(models.APIResponse{Data: tour})
}
