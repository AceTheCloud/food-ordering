package main

import (
	"log"

	"github.com/acethecloud/food-ordering/database"
	"github.com/acethecloud/food-ordering/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()

	app := fiber.New()

	routes.CreateRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
