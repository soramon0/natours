package main

import (
	"natours/pkg/configs"
	"natours/pkg/database"
	"natours/pkg/models"
	"natours/pkg/routes"
	"natours/pkg/utils"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client := database.OpenConnection()
	defer database.CloseConnection(client)

	app := fiber.New(configs.FiberConfig())
	services := models.NewServices(client)
	logger := utils.InitLogger()

	routes.Register(app, services, logger)
	utils.StartServer(app, logger)
}
