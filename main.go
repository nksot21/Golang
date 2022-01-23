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

func main() {
	app := fiber.New()

	driver.Connect(host, user, password, dbname, port)

	// Set up routes
	routes.SetUpUserRoutes(app)
	routes.SetUpProductRoutes(app)
	routes.SetUpOrderRoutes(app)
	routes.SetUpSiteRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
