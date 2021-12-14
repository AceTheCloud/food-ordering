package handlers

import (
	"errors"

	"github.com/acethecloud/food-ordering/database"
	"github.com/acethecloud/food-ordering/models"
	"github.com/gofiber/fiber/v2"
)

type Dish struct {
	// This is not the model, more like a serializer
	ID         uint `json:"id"`
	Name       string
	Available  bool              `gorm:"default:true"`
	Price      uint              `json:"price"`
	Restaurant models.Restaurant `json:"restaurant" gorm:"foreignKey:RestaurantRefer"`
	Cuisine    models.Cuisine    `json:"Cuisine" gorm:"foreignKey:CuisineRefer"`
}

func CreateDish(ctx *fiber.Ctx) error {
	var dish models.Dish

	if err := ctx.BodyParser(&dish); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var restaurant models.Restaurant

	if err := findRestaurant(dish.RestaurantRefer, &restaurant); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var cuisine models.Cuisine

	if err := findCuisine(dish.CuisineRefer, &cuisine); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&dish)
	dish.Cuisine = cuisine
	dish.Restaurant = restaurant

	responseDish := createResponseDish(dish)
	return ctx.Status(200).JSON(responseDish)
}

func createResponseDish(dish models.Dish) Dish {
	return Dish{
		ID:         dish.ID,
		Name:       dish.Name,
		Available:  dish.Available,
		Price:      dish.Price,
		Restaurant: dish.Restaurant,
		Cuisine:    dish.Cuisine,
	}
}

func GetDishesByRestaurant(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var restaurant models.Restaurant
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findRestaurant(uint(id), &restaurant); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	dishes := []models.Dish{}

	var cuisines []models.Cuisine
	cuisines, err = getAllCuisines()

	if err != nil {
		return c.Status(400).JSON("No Cuisines found")
	}

	database.Database.Db.Where("deleted = ? AND restaurant_refer = ?", false, restaurant.ID).Find(&dishes)
	responseDishes := []Dish{}

	for _, dish := range dishes {
		dish.Restaurant = restaurant
		dish.Cuisine = findCuisineById(cuisines, dish.CuisineRefer)
		responseDish := createResponseDish(dish)
		responseDishes = append(responseDishes, responseDish)
	}
	return c.Status(200).JSON(responseDishes)
}

func findCuisineById(cuisines []models.Cuisine, id uint) models.Cuisine {
	for _, cuisine := range cuisines {
		if cuisine.ID == id {
			return cuisine
		}
	}
	return models.Cuisine{}
}

func findDish(id int, dish *models.Dish) error {
	database.Database.Db.Find(&dish, "deleted = ? AND id = ?", false, id)
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

	var restaurant models.Restaurant
	if err := findRestaurant(uint(dish.RestaurantRefer), &restaurant); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	dish.Restaurant = restaurant

	var cuisine models.Cuisine
	if err := findCuisine(uint(dish.CuisineRefer), &cuisine); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	dish.Cuisine = cuisine

	responseDish := createResponseDish(dish)
	return c.Status(200).JSON(responseDish)
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
		Price     uint
	}

	var updateData UpdateDish

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	dish.Name = updateData.Name
	dish.Available = updateData.Available
	dish.Price = updateData.Price

	database.Database.Db.Save(&dish)

	var restaurant models.Restaurant
	if err := findRestaurant(uint(dish.RestaurantRefer), &restaurant); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	dish.Restaurant = restaurant

	var cuisine models.Cuisine
	if err := findCuisine(uint(dish.CuisineRefer), &cuisine); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	dish.Cuisine = cuisine

	responseDish := createResponseDish(dish)

	return c.Status(200).JSON(responseDish)
}

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
	dish.Deleted = true

	//Updating to showcase softdelete
	//In case of real delete simply use Db.Delete(&dish)
	if err = database.Database.Db.Save(&dish).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Dish")
}
