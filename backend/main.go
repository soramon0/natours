package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/soramon0/natrous/pkg/configs"
	"github.com/soramon0/natrous/pkg/database"
	"github.com/soramon0/natrous/pkg/models"
	"github.com/soramon0/natrous/pkg/routes"
	"github.com/soramon0/natrous/pkg/utils"
)

func main() {
	app := fiber.New(configs.FiberConfig())
	client := database.OpenConnection()
	defer database.CloseConnection(client)

	services := models.NewServices(client)
	logger := utils.InitLogger()

	routes.Register(app, services, logger)
	utils.StartServer(app)
}
