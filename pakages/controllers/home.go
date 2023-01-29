package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func (controller Controller) HomeController(context *fiber.Ctx) error {
	return context.JSON(&fiber.Map{
		"success": true,
		"content": "Hello World with Fiber!!",
	})
}
