package routes

import (
	"log"

	"natours/pkg/middleware"
	"natours/pkg/models"
	"natours/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(a *fiber.App, s *models.Services, vt *utils.ValidatorTransaltor, l *log.Logger) {
	middleware.FiberMiddleware(a)

	registerUserRoutes(a, s, l)
	registerTourRoutes(a, s, vt, l)
	registerNotFoundRoute(a)
}
