package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/controllers"
	"github.com/soramon0/natrous/pkg/models"
)

func registerUserRoutes(a *fiber.App, s *models.Services, l *log.Logger) *fiber.Router {
	router := a.Group("/api/v1/users")
	usersC := controllers.NewUsers(s.User, l)

	router.Get("/", usersC.GetUsers)
	router.Get("/:id", usersC.GetUser)
	router.Post("/", usersC.CreateUser)

	return &router
}
