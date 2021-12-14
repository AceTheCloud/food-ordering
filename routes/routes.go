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
	app.Delete("/api/users/:id", handlers.DeleteUser)

	// Restaurants endpoints
	app.Post("/api/restaurants", handlers.CreateRestaurant)
	app.Get("/api/restaurants", handlers.GetRestaurants)
	app.Get("/api/restaurants/:id", handlers.GetRestaurant)
	app.Put("/api/restaurants/:id", handlers.UpdateRestaurant)
	app.Delete("/api/restaurants/:id", handlers.DeleteRestaurant)

	// Cuisines endpoints
	app.Post("/api/cuisines", handlers.CreateCuisine)
	app.Get("/api/cuisines", handlers.GetCuisines)
	app.Get("/api/cuisines/:id", handlers.GetCuisine)
	app.Put("/api/cuisines/:id", handlers.UpdateCuisine)
	app.Delete("/api/cuisines/:id", handlers.DeleteCuisine)

	// Cuisines endpoints
	app.Post("/api/dish", handlers.CreateDish)
	app.Get("/api/dishesByRestaurant/:id", handlers.GetDishesByRestaurant)
	app.Get("/api/dish/:id", handlers.GetDish)
	app.Put("/api/dish/:id", handlers.UpdateDish)
	app.Delete("/api/dish/:id", handlers.DeleteDish)
}
