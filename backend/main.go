package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	Message    string `json:"message"`
	statusCode int
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
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code and message defaults to 500
			apiError := APIResponse{Error: &APIError{
				Message:    fiber.ErrInternalServerError.Error(),
				statusCode: fiber.StatusInternalServerError,
			}}

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				apiError.Error.statusCode = e.Code
				apiError.Error.Message = e.Message
			}

			// Send custom error response
			if err := ctx.Status(apiError.Error.statusCode).JSON(apiError); err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(APIResponse{Error: &APIError{
					Message: "Internal Server Error",
				}})
			}

			// Return from handler
			return nil
		},
	})
	app.Use(recover.New())
	app.Use(logger.New())

	toursApi := app.Group("/api/v1/tours")

	toursApi.Get("/", func(c *fiber.Ctx) error {
		var tours []Tour
		if err := readJsonFile("tours-simple", &tours); err != nil {
			return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
		}

		return c.JSON(APIResponse{Data: tours, Count: len(tours)})
	})

	toursApi.Get("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid Id"}
		}

		var tours []Tour
		if err := readJsonFile("tours-simple", &tours); err != nil {
			return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
		}

		var tour *Tour
		for _, item := range tours {
			if item.Id == id {
				tour = &item
				break
			}
		}

		if tour == nil {
			return &fiber.Error{Code: fiber.StatusNotFound, Message: "Tour not found"}
		}

		return c.JSON(APIResponse{Data: tour, Count: 1})
	})

	toursApi.Post("/", func(c *fiber.Ctx) error {
		payload := struct {
			Name string `json:"name"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
		}

		if payload.Name == "" {
			return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Name is required"}
		}

		var tours []Tour
		if err := readJsonFile("tours-simple", &tours); err != nil {
			return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
		}

		tour := Tour{Name: payload.Name, Id: len(tours)}
		tours = append(tours, tour)
		if err := writeJsonFile("tours-simple", &tours); err != nil {
			return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
		}

		return c.Status(fiber.StatusCreated).JSON(APIResponse{Data: tour, Count: 1})
	})

	toursApi.Patch("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid Id"}
		}

		payload := Tour{}
		if err := c.BodyParser(&payload); err != nil {
			return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
		}

		var tours []Tour
		if err := readJsonFile("tours-simple", &tours); err != nil {
			return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
		}

		index := -1
		for i, item := range tours {
			if item.Id == id {
				index = i
				break
			}
		}

		if index == -1 {
			return &fiber.Error{Code: fiber.StatusNotFound, Message: "Tour not found"}
		}

		payload.Id = id
		tours[index] = payload
		if err := writeJsonFile("tours-simple", &tours); err != nil {
			return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
		}

		return c.Status(fiber.StatusCreated).JSON(APIResponse{Data: tours[index], Count: 1})
	})

	log.Fatalln(app.Listen(":5000"))
}
