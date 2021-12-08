package handlers

import (
	"errors"

	"github.com/acethecloud/food-ordering/database"
	"github.com/acethecloud/food-ordering/models"
	"github.com/gofiber/fiber/v2"
)

type Cuisine struct {
	// This is not the model, more like a serializer
	ID   uint `json:"id"`
	Name string
}

func CreateCuisine(ctx *fiber.Ctx) error {
	var cuisine models.Cuisine

	if err := ctx.BodyParser(&cuisine); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&cuisine)
	responseCuisine := createResponseCuisine(cuisine)
	return ctx.Status(200).JSON(responseCuisine)
}

func createResponseCuisine(cuisine models.Cuisine) Cuisine {
	return Cuisine{
		ID:   cuisine.ID,
		Name: cuisine.Name,
	}
}

func GetCuisines(c *fiber.Ctx) error {
	cuisines := []models.Cuisine{}
	database.Database.Db.Find(&cuisines)
	responseCuisines := []Cuisine{}

	for _, cuisine := range cuisines {
		responseCuisine := createResponseCuisine(cuisine)
		responseCuisines = append(responseCuisines, responseCuisine)
	}
	return c.Status(200).JSON(responseCuisines)
}

func findCuisine(id int, cuisine *models.Cuisine) error {
	database.Database.Db.Find(&cuisine, "id = ?", id)
	if cuisine.ID == 0 {
		return errors.New("Cuisine does not exist")
	}
	return nil
}

func GetCuisine(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var cuisine models.Cuisine
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findCuisine(id, &cuisine); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseCuisine := createResponseCuisine(cuisine)
	return c.Status(200).JSON(responseCuisine)
}

func UpdateCuisine(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var cuisine models.Cuisine
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findCuisine(id, &cuisine)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateCuisine struct {
		Name string
	}

	var updateData UpdateCuisine

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	cuisine.Name = updateData.Name

	database.Database.Db.Save(&cuisine)

	responseCuisine := createResponseCuisine(cuisine)

	return c.Status(200).JSON(responseCuisine)
}

// Should be used only by root user or admin users
func DeleteCuisine(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var cuisine models.Cuisine

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findCuisine(id, &cuisine)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&cuisine).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Cuisine")
}
