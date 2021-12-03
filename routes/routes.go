package routes

import (
	"github.com/acethecloud/food-ordering/handlers"
	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(app *fiber.App) {
	//Health endpoint
	app.Get("/api/health", handlers.CheckHealth)

	// User endpoints
	app.Post("/api/users", handlers.CreateUser)
	app.Get("/api/users", handlers.GetUsers)
	app.Get("/api/users/:id", handlers.GetUser)
	app.Delete("/api/users/:id", handlers.DeleteUser)
}
