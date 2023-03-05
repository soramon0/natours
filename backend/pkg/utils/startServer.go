package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	url := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)

	// Run server.
	if err := a.Listen(url); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v\n", err)
	}
}
