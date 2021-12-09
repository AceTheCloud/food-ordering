package handlers

import (
	"github.com/acethecloud/food-ordering/database"
	"github.com/acethecloud/food-ordering/enums"
	"github.com/acethecloud/food-ordering/models"
	"github.com/gofiber/fiber/v2"
)

type OrderItemResponse struct {
	ID uint `json:"id"`

	OrderId string `json:"order_id"`

	DishId uint `json:"dish_id"`

	Quantity uint `json:"quantity"`
}

type OrderResponse struct {
	ID uint `json:"id"`

	User User `json:"user"`

	TotalPrice uint `json:"total_price"`

	Status enums.Status `json:"status"`
}

func CreateResponseOrder(order models.Order, user User) OrderResponse {
	return OrderResponse{
		ID:         order.ID,
		User:       user,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}
}

func CreateResponseOrderItem(orderItem models.OrderItem) OrderItemResponse {
	return OrderItemResponse{
		ID:       orderItem.ID,
		OrderId:  orderItem.OrderId,
		DishId:   orderItem.DishId,
		Quantity: orderItem.Quantity,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	if err := findUser(uint(order.UserRefer), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseOrder := CreateResponseOrder(order, responseUser)

	return c.Status(200).JSON(responseOrder)
}

func CreateOrderItem(c *fiber.Ctx) error {
	var orderItem models.OrderItem

	if err := c.BodyParser(&orderItem); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&orderItem)

	responseOrder := CreateResponseOrderItem(orderItem)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []OrderResponse{}

	for _, order := range orders {
		var user models.User
		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)
}
