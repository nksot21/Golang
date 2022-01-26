package main

import (
	"connectdb/src/driver"
	"connectdb/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// go get golang.org/x/crypto/bcrypt
// jwtware "github.com/gofiber/jwt/v3"
const (
	host     = "localhost"
	user     = "postgres"
	port     = "5432"
	password = "Nhakhuyen21."
	dbname   = "memories"
)

func SetUpRoutes(app *fiber.App) {
	// USER ROUTES
	user := app.Group("/user")
	user.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	routes.SetUpUserRoutes(user)

	// PRODUCT ROUTES
	product := app.Group("/product")
	product.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	routes.SetUpProductRoutes(product)

	// ORDER ROUTES
	order := app.Group("/order")
	order.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	routes.SetUpOrderRoutes(order)

	// SITE ROUTES
	routes.SetUpSiteRoutes(app)

	// AUTHOR ROUTES
	routes.SetUpAuthorRoutes(app)
}

func main() {
	app := fiber.New()

	driver.Connect(host, user, password, dbname, port)

	SetUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
