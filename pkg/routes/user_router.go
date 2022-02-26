package routes

import (
	"github.com/gofiber/fiber/v2"
	"mental-health-api/hanlder"
)

func UserRouter(a *fiber.App) {
	router := a.Group("/user")
	router.Post("/", hanlder.CreateUser)
}
