package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/natrous/pkg/models"
)

func GetTours(c *fiber.Ctx) error {
	var tours []models.Tour
	if err := readJsonFile("tours-simple", &tours); err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(models.APIResponse{Data: tours, Count: len(tours)})
}

func GetTour(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid Id"}
	}

	var tours []models.Tour
	if err := readJsonFile("tours-simple", &tours); err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	var tour *models.Tour
	for _, item := range tours {
		if item.Id == id {
			tour = &item
			break
		}
	}

	if tour == nil {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: "Tour not found"}
	}

	return c.JSON(models.APIResponse{Data: tour, Count: 1})
}

func CreateTour(c *fiber.Ctx) error {
	payload := struct {
		Name string `json:"name"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	if payload.Name == "" {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Name is required"}
	}

	var tours []models.Tour
	if err := readJsonFile("tours-simple", &tours); err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	tour := models.Tour{Name: payload.Name, Id: len(tours)}
	tours = append(tours, tour)
	if err := writeJsonFile("tours-simple", &tours); err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	return c.Status(fiber.StatusCreated).JSON(models.APIResponse{Data: tour, Count: 1})
}

func UpdateTour(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "Invalid Id"}
	}

	payload := models.Tour{}
	if err := c.BodyParser(&payload); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	var tours []models.Tour
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

	return c.Status(fiber.StatusCreated).JSON(models.APIResponse{Data: tours[index], Count: 1})
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
