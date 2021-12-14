package handlers

import (
	"errors"

	"github.com/acethecloud/food-ordering/database"
	"github.com/acethecloud/food-ordering/models"
	"github.com/gofiber/fiber/v2"
)

type Restaurant struct {
	// This is not the model, more like a serializer
	ID          uint `json:"id"`
	Name        string
	City        string
	OpeningTime uint8 `json:"opening_time"`
	ClosingTime uint8 `json:"closing_time"`
	Open        bool
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
		ID:          restaurant.ID,
		Name:        restaurant.Name,
		City:        restaurant.City,
		OpeningTime: restaurant.OpeningTime,
		ClosingTime: restaurant.ClosingTime,
		Open:        restaurant.Open,
	}
}

func GetRestaurants(c *fiber.Ctx) error {
	restaurants := []models.Restaurant{}
	database.Database.Db.Where("deleted", false).Find(&restaurants)
	responseRestaurants := []Restaurant{}
	for _, restaurant := range restaurants {
		responseRestaurant := createResponseRestaurant(restaurant)
		responseRestaurants = append(responseRestaurants, responseRestaurant)
	}

	return c.Status(200).JSON(responseRestaurants)
}

func findRestaurant(id uint, restaurant *models.Restaurant) error {
	database.Database.Db.Find(&restaurant, "deleted = ? AND id = ?", false, id)
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

	if err := findRestaurant(uint(id), &restaurant); err != nil {
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

	err = findRestaurant(uint(id), &restaurant)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateRestaurant struct {
		Name        string
		City        string
		Open        bool
		OpeningTime uint8 `json:"opening_time"`
		ClosingTime uint8 `json:"closing_time"`
	}

	var updateData UpdateRestaurant

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	restaurant.Name = updateData.Name
	restaurant.City = updateData.City
	restaurant.Open = updateData.Open
	restaurant.OpeningTime = updateData.OpeningTime
	restaurant.ClosingTime = updateData.ClosingTime

	database.Database.Db.Save(&restaurant)

	responseRestaurant := createResponseRestaurant(restaurant)

	return c.Status(200).JSON(responseRestaurant)
}

// Should be used only by root user or admin users
func DeleteRestaurant(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var restaurant models.Restaurant

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findRestaurant(uint(id), &restaurant)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	restaurant.Deleted = true

	//Updating to showcase softdelete
	//In case of real delete simply use Db.Delete(&restaurant)
	if err = database.Database.Db.Save(&restaurant).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Restaurant")
}
