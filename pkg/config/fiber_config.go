package configs

import (
	"github.com/gofiber/fiber/v2"
	models "mental-health-api/model"
	"os"
	"strconv"
	"time"
)

func FiberConfig() fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		ErrorHandler: fiberErrorHandler,
	}
}

func fiberErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	err = ctx.Status(code).JSON(models.Response{
		Status:  code,
		Message: err.Error(),
		Error:   true,
		Data:    nil,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal server error",
			Error:   true,
			Data:    nil,
		})
	}

	return nil
}
