package routes

import (
	"connectdb/src/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpProductRoutes(app *fiber.App) {
	// POST: /newproduct
	app.Post("/newproduct", handler.CreateProduct)
	// GET: /product
	app.Get("/product", handler.GetProducts)
	// GET: /product/:id
	app.Get("/product/:id", handler.GetProduct)
	// PUT: /product/:id
	app.Put("/product/:id", handler.UpdateProduct)
	// DELETE: /product/:id
	app.Delete("/product/:id", handler.DeleteProduct)
}
