package main

import (
	"connectdb/src/driver"
	"connectdb/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	host     = "localhost"
	user     = "postgres"
	port     = "5432"
	password = "Nhakhuyen21."
	dbname   = "memories"
)

func setUpRoutes(app *fiber.App) {
	// GET: /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome")
	})
	// POST: /newuser
	app.Post("/newuser", routes.CreateUser)
	// GET: /user
	app.Get("/user", routes.GetUsers)
	// GET: /user/:id
	app.Get("/user/:id", routes.GetUser)
	// PUT: /user/:id
	app.Put("/user/:id", routes.UpdateUser)
	// DELETE: /user/:id
	app.Delete("/user/:id", routes.DeleteUser)

	// POST: /newproduct
	app.Post("/newproduct", routes.CreateProduct)
	// GET: /product
	app.Get("/product", routes.GetProducts)
	// GET: /product/:id
	app.Get("/product/:id", routes.GetProduct)
	// PUT: /product/:id
	app.Put("/product/:id", routes.UpdateProduct)
	// DELETE: /product/:id
	app.Delete("/product/:id", routes.DeleteProduct)
}

func main() {
	app := fiber.New()

	driver.Connect(host, user, password, dbname, port)
	setUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
