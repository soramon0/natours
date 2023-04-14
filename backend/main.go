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
	logger := utils.InitLogger()
	client := database.OpenConnection(utils.GetDatabaseBindAdress(), logger)
	defer func() {
		utils.Must(database.CloseConnection(client))
	}()
	if err := database.CreateIndexes(client); err != nil {
		logger.Fatalln("could not create db indexes", err)
	}

	app := fiber.New(configs.FiberConfig())
	services := models.NewServices(client)

	vt, err := utils.NewValidator()
	if err != nil {
		logger.Fatalf("could not create validator %v\n", err)
	}

	routes.Register(app, services, vt, logger)
	utils.StartServer(app, logger)
}
