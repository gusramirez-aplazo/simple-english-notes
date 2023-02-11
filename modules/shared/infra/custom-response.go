package infra

import (
	"github.com/gofiber/fiber/v2"
)

func CustomResponse(
	context *fiber.Ctx,
	status int,
	isSuccess bool,
	content any,
	message string,
) error {
	return context.
		Status(status).
		JSON(fiber.Map{
			"success": isSuccess,
			"content": content,
			"message": message,
		})
}
