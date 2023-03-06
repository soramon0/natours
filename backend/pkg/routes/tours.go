package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/controllers"
	"github.com/soramon0/natrous/pkg/models"
)

func registerTourRoutes(a *fiber.App, s *models.Services, l *log.Logger) *fiber.Router {
	// Create routes group.
	router := a.Group("/api/v1/tours")

	router.Get("/", controllers.GetTours)
	router.Get("/:id", controllers.GetTour)
	router.Post("/", controllers.CreateTour)
	router.Patch("/:id", controllers.UpdateTour)

	return &router
}
