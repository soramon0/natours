package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Tour struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type APIResponse struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
	Error *APIError   `json:"error"`
}

type APIError struct {
	Message string `json:"message"`
}

func readJsonData(filename string) *[]Tour {
	data, err := os.ReadFile(fmt.Sprintf("./data/%s.json", filename))
	if err != nil {
		log.Fatalln(err)
	}

	var tours []Tour

	if err = json.Unmarshal(data, &tours); err != nil {
		log.Fatalln(err)
	}

	return &tours
}

func main() {
	tours := readJsonData("tours-simple")
	app := fiber.New()
	toursApi := app.Group("/api/v1/tours")

	toursApi.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(APIResponse{Data: tours, Count: len(*tours)})
	})
	toursApi.Post("/", func(c *fiber.Ctx) error {
		return c.JSON(APIResponse{Error: &APIError{Message: fiber.ErrNotImplemented.Message}})
	})

	log.Fatalln(app.Listen(":5000"))
}
