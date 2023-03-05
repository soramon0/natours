package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/soramon0/natrous/pkg/configs"
	"github.com/soramon0/natrous/pkg/database"
	"github.com/soramon0/natrous/pkg/middleware"
	"github.com/soramon0/natrous/pkg/routes"
	"github.com/soramon0/natrous/pkg/utils"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	database.OpenConnection()
	defer database.CloseConnection()

	middleware.FiberMiddleware(app)
	routes.UserRoutes(app)
	routes.TourRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServer(app)
}
