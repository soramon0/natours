package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/controllers"
)

func TourRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1/tours")

	route.Get("/", controllers.GetTours)
	route.Get("/:id", controllers.GetTour)
	route.Post("/", controllers.CreateTour)
	route.Patch("/:id", controllers.UpdateTour)
}
