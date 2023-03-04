package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/configs"
	"github.com/soramon0/natrous/pkg/middleware"
	"github.com/soramon0/natrous/pkg/routes"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	middleware.FiberMiddleware(app)
	routes.TourRoutes(app)

	log.Fatalln(app.Listen(":5000"))
}
