package routes

import (
	"connectdb/src/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpOrderRoutes(order fiber.Router) {
	// POST: order/new
	order.Post("/", handler.CreateOrder)
	// GET: /order
	order.Get("/", handler.GetOrders)
	// GET: /order/:id
	order.Get("/:id", handler.GetOrder)
}
