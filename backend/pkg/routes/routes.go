package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/middleware"
	"github.com/soramon0/natrous/pkg/models"
)

func Register(a *fiber.App, s *models.Services, l *log.Logger) {
	middleware.FiberMiddleware(a)

	registerUserRoutes(a, s, l)
	registerTourRoutes(a, s, l)
	registerNotFoundRoute(a)
}
