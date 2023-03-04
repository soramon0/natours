package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Tour struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Duration        int      `json:"duration"`
	MaxGroupSize    int      `json:"maxGroupSize"`
	Difficulty      string   `json:"difficulty"`
	RatingsAverage  float64  `json:"ratingsAverage"`
	RatingsQuantity int      `json:"ratingsQuantity"`
	Price           int      `json:"price"`
	Summary         string   `json:"summary"`
	Description     string   `json:"description"`
	ImageCover      string   `json:"imageCover"`
	Images          []string `json:"images"`
	StartDates      []string `json:"startDates"`
}

type APIResponse struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
	Error *APIError   `json:"error"`
}

type APIError struct {
	Message string `json:"message"`
}

func readJsonFile(filename string, data any) error {
	jsonFile, err := os.ReadFile(fmt.Sprintf("./data/%s.json", filename))
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonFile, &data)
}

func writeJsonFile(filename string, data any) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("./data/%s.json", filename), file, 0644)
}

func main() {
	app := fiber.New()
	toursApi := app.Group("/api/v1/tours")

	toursApi.Get("/", func(c *fiber.Ctx) error {
		var tours []Tour
		if err := readJsonFile("tours-simple", &tours); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{Error: &APIError{Message: err.Error()}})
		}

		return c.JSON(APIResponse{Data: tours, Count: len(tours)})
	})

	toursApi.Post("/", func(c *fiber.Ctx) error {
		payload := struct {
			Name string `json:"name"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(APIResponse{Error: &APIError{Message: err.Error()}})
		}

		if payload.Name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(APIResponse{Error: &APIError{Message: "Name is required"}})
		}

		var tours []Tour
		if err := readJsonFile("tours-simple", &tours); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{Error: &APIError{Message: err.Error()}})
		}

		tour := Tour{Name: payload.Name, Id: len(tours)}
		tours = append(tours, tour)
		if err := writeJsonFile("tours-simple", &tours); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{Error: &APIError{Message: err.Error()}})
		}

		return c.Status(fiber.StatusCreated).JSON(APIResponse{Data: tour})
	})

	log.Fatalln(app.Listen(":5000"))
}
