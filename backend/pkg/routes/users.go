package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/controllers"
)

func UserRoutes(a *fiber.App) {
	route := a.Group("/api/v1/users")

	route.Get("/", controllers.GetUsers)
	route.Get("/:id", controllers.GetUser)
	route.Post("/", controllers.CreateUser)
}
