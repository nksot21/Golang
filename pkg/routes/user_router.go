package routes

import (
	"github.com/gofiber/fiber/v2"
	"mental-health-api/hanlder"
)

func UserRouter(a *fiber.App) {
	router := a.Group("/user")

	router.Get("/", hanlder.GetUser)
	router.Post("/", hanlder.CreateUser)
	router.Put("/", hanlder.UpdateUser)

	router.Post("/feel", hanlder.AddFeelUser)
	//router.Get("/feel", hanlder.GetFeelUser)
}
