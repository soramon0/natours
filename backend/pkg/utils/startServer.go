package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App, l *log.Logger) {

	// Run server.
	if err := a.Listen(GetServerBindAddress()); err != nil {
		l.Printf("Oops... Server is not running! Reason: %v\n", err)
	}
}
