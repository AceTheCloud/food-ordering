package routes

import (
	"github.com/acethecloud/food-ordering/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	//Health endpoint
	app.Get("/api/health", handlers.CheckHealth)

	// User endpoints
	app.Post("/api/users", handlers.CreateUser)
	app.Get("/api/users", handlers.GetUsers)
	app.Get("/api/users/:id", handlers.GetUser)
	app.Put("/api/users/:id", handlers.UpdateUser)

	// Restaurants endpoints
	app.Post("/api/restaurants", handlers.CreateRestaurant)
	app.Get("/api/restaurants", handlers.GetRestaurants)
	app.Get("/api/restaurants/:id", handlers.GetRestaurant)
	app.Put("/api/restaurants/:id", handlers.UpdateRestaurant)

	// Cuisines endpoints
	app.Post("/api/cuisines", handlers.CreateCuisine)
	app.Get("/api/cuisines", handlers.GetCuisines)
	app.Get("/api/cuisines/:id", handlers.GetCuisine)
	app.Put("/api/cuisines/:id", handlers.UpdateCuisine)

	// Cuisines endpoints
	app.Post("/api/cuisines", handlers.CreateCuisine)
	app.Get("/api/cuisines", handlers.GetCuisines)
	app.Get("/api/cuisines/:id", handlers.GetCuisine)
	app.Put("/api/cuisines/:id", handlers.UpdateCuisine)
}
