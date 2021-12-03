package handlers

import (
	"errors"

	"github.com/acethecloud/food-ordering/database"
	"github.com/acethecloud/food-ordering/models"
	"github.com/gofiber/fiber/v2"
)

type Restaurant struct {
	// This is not the model, more like a serializer
	ID       uint `json:"id"`
	Name     string
	City     string
	Cuisines []string
}

func CreateRestaurant(ctx *fiber.Ctx) error {
	var restaurant models.Restaurant

	if err := ctx.BodyParser(&restaurant); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&restaurant)
	responseRestaurant := createResponseRestaurant(restaurant)
	return ctx.Status(200).JSON(responseRestaurant)
}

func createResponseRestaurant(restaurant models.Restaurant) Restaurant {
	return Restaurant{
		ID:       restaurant.ID,
		Name:     restaurant.Name,
		City:     restaurant.City,
		Cuisines: restaurant.Cuisines,
	}
}

func GetRestaurants(c *fiber.Ctx) error {
	restaurants := []models.Restaurant{}
	database.Database.Db.Find(&restaurants)
	responseRestaurants := []Restaurant{}
	for _, restaurant := range restaurants {
		responseRestaurant := createResponseRestaurant(restaurant)
		responseRestaurants = append(responseRestaurants, responseRestaurant)
	}

	return c.Status(200).JSON(responseRestaurants)
}

func findRestaurant(id int, restaurant *models.Restaurant) error {
	database.Database.Db.Find(&restaurant, "id = ?", id)
	if restaurant.ID == 0 {
		return errors.New("Restaurant does not exist")
	}
	return nil
}

func GetRestaurant(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var restaurant models.Restaurant

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findRestaurant(id, &restaurant); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseRestaurant := createResponseRestaurant(restaurant)

	return c.Status(200).JSON(responseRestaurant)
}

func UpdateRestaurant(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var restaurant models.Restaurant

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findRestaurant(id, &restaurant)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateRestaurant struct {
		Name     string
		City     string
		Cuisines []string
	}

	var updateData UpdateRestaurant

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	restaurant.Name = updateData.Name
	restaurant.City = updateData.City
	restaurant.Cuisines = updateData.Cuisines

	database.Database.Db.Save(&restaurant)

	responseRestaurant := createResponseRestaurant(restaurant)

	return c.Status(200).JSON(responseRestaurant)
}

func DeleteRestaurant(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var restaurant models.Restaurant

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findRestaurant(id, &restaurant)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&restaurant).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Restaurant")
}
