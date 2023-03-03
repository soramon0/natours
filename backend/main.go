package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Tour struct {
	Name string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(Tour{Name: "Hello"})
	})

	log.Fatalln(app.Listen(":3000"))
}
