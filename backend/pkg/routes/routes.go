package routes

import (
	"log"

	"natours/pkg/middleware"
	"natours/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func Register(a *fiber.App, s *models.Services, l *log.Logger) {
	middleware.FiberMiddleware(a)

	registerUserRoutes(a, s, l)
	registerTourRoutes(a, s, l)
	registerNotFoundRoute(a)
}
