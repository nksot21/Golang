package routes

import (
	"connectdb/src/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpProductRoutes(product fiber.Router) {
	// POST: order/new
	product.Post("/", handler.CreateProduct)
	// GET: /product
	product.Get("/", handler.GetProducts)
	// GET: /product/:id
	product.Get("/:id", handler.GetProduct)
	// PUT: /product/:id
	product.Put("/:id", handler.UpdateProduct)
	// DELETE: /product/:id
	product.Delete("/:id", handler.DeleteProduct)
}
