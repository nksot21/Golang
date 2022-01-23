package routes

import (
	"connectdb/src/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpOrderRoutes(app *fiber.App) {
	// POST: /neworder
	app.Post("/neworder", handler.CreateOrder)
	// GET: /order
	app.Get("/order", handler.GetOrders)
	// GET: /order/:id
	app.Get("/order/:id", handler.GetOrder)
}
