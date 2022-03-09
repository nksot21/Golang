package routes

import (
	"mental-health-api/handler"

	"github.com/gofiber/fiber/v2"
)

func UserFeelRouter(a *fiber.App) {
	router := a.Group("/user-feel")
	router.Get("/", handler.GetUserFeel)
	router.Post("/", handler.CreateUserFeel)
}
