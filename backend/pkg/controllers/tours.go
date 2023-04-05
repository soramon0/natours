package controllers

import (
	"errors"
	"log"

	"natours/pkg/models"
	"natours/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tours struct {
	service models.TourService
	vt      *utils.ValidatorTransaltor
	log     *log.Logger
}

// New Users is used to create a new Users controller.
func NewTours(ts models.TourService, vt *utils.ValidatorTransaltor, l *log.Logger) *Tours {
	return &Tours{
		service: ts,
		vt:      vt,
		log:     l,
	}
}

func (t *Tours) GetTours(c *fiber.Ctx) error {
	tours, err := t.service.Find()
	if err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(models.NewAPIResponse(tours, len(tours)))
}

func (t *Tours) GetTour(c *fiber.Ctx) error {
	tour, err := t.service.ByID(c.Params("id"))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &fiber.Error{Code: fiber.StatusNotFound, Message: "Tour not found"}
		}

		t.log.Println(err)
		return &fiber.Error{Code: fiber.StatusNotFound, Message: err.Error()}
	}

	return c.JSON(models.NewAPIResponse(tour, 0))
}

func (t *Tours) CreateTour(c *fiber.Ctx) error {
	var payload models.CreateTourPayload
	if err := c.BodyParser(&payload); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}
	if err := t.vt.Validator.Struct(&payload); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return c.Status(fiber.StatusBadRequest).JSON(t.vt.ValidationErrors(ve))
		}
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	tour, err := t.service.Create(payload)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	return c.JSON(models.NewAPIResponse(tour, 0))
}

func (t *Tours) UpdateTour(c *fiber.Ctx) error {
	tour, err := t.service.ByID(c.Params("id"))
	if err != nil {
		t.log.Println(err)
		return &fiber.Error{Code: fiber.StatusNotFound, Message: err.Error()}
	}

	payload := models.Tour{}
	if err := c.BodyParser(&payload); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	return c.Status(fiber.StatusCreated).JSON(models.NewAPIResponse(tour, 0))
}
