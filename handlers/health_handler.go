package handlers

import "github.com/gofiber/fiber/v2"

func CheckHealth(ctx *fiber.Ctx) error {
	return ctx.SendString("Healthy")
}
