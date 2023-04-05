package routes

import (
	"log"

	"natours/pkg/controllers"
	"natours/pkg/models"
	"natours/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func registerTourRoutes(a *fiber.App, s *models.Services, vt *utils.ValidatorTransaltor, l *log.Logger) *fiber.Router {
	router := a.Group("/api/v1/tours")
	toursC := controllers.NewTours(s.Tour, vt, l)

	router.Get("/", toursC.GetTours)
	router.Get("/:id", toursC.GetTour)
	router.Post("/", toursC.CreateTour)
	router.Patch("/:id", toursC.UpdateTour)

	return &router
}
