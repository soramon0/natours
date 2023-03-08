package routes

import (
	"log"

	"natours/pkg/controllers"
	"natours/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func registerTourRoutes(a *fiber.App, s *models.Services, l *log.Logger) *fiber.Router {
	router := a.Group("/api/v1/tours")
	toursC := controllers.NewTours(s.Tour, l)

	router.Get("/", toursC.GetTours)
	router.Get("/:id", toursC.GetTour)
	router.Post("/", toursC.CreateTour)
	router.Patch("/:id", toursC.UpdateTour)

	return &router
}
