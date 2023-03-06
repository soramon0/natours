package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/controllers"
	"github.com/soramon0/natrous/pkg/models"
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
