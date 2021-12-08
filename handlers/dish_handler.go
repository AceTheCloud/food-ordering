package handlers

import (
	"errors"

	"github.com/acethecloud/food-ordering/database"
	"github.com/acethecloud/food-ordering/models"
	"github.com/gofiber/fiber/v2"
)

type Dish struct {
	// This is not the model, more like a serializer
	ID        uint `json:"id"`
	Name      string
	Available bool `gorm:"default:true"`
	Deleted   bool `gorm:"default:false"`
}

func CreateDish(ctx *fiber.Ctx) error {
	var dish models.Dish

	if err := ctx.BodyParser(&dish); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&dish)
	responseDish := createResponseDish(dish)
	return ctx.Status(200).JSON(responseDish)
}

func createResponseDish(dish models.Dish) Dish {
	return Dish{
		ID:        dish.ID,
		Name:      dish.Name,
		Available: dish.Available,
	}
}

func GetDishes(c *fiber.Ctx) error {
	dishes := []models.Dish{}
	database.Database.Db.Find(&dishes)
	responseDishes := []Dish{}

	for _, dish := range dishes {
		responseDish := createResponseDish(dish)
		responseDishes = append(responseDishes, responseDish)
	}
	return c.Status(200).JSON(responseDishes)
}

func findDish(id int, dish *models.Dish) error {
	database.Database.Db.Find(&dish, "id = ?", id)
	if dish.ID == 0 {
		return errors.New("Dish does not exist")
	}
	return nil
}

func GetDish(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var dish models.Dish
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findDish(id, &dish); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseDish := createResponseDish(dish)
	return c.Status(200).JSON(responseDish.ID)
}

func UpdateDish(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var dish models.Dish
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findDish(id, &dish)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateDish struct {
		Name      string
		Available bool
	}

	var updateData UpdateDish

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	dish.Name = updateData.Name
	dish.Available = updateData.Available

	database.Database.Db.Save(&dish)

	responseDish := createResponseDish(dish)

	return c.Status(200).JSON(responseDish)
}

// Should be used only by root user or admin users
func DeleteDish(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var dish models.Dish

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findDish(id, &dish)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&dish).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Dish")
}
