package main

import (
	"DBdemo/src/driver"
	"DBdemo/src/model"
	"DBdemo/src/repository"
	"DBdemo/src/repository/implement"
	"DBdemo/src/router"

	//"DBdemo/src/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// database connection constant
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "Nhakhuyen21."
	dbname   = "memories"
)

func setUpRoutes(clientRepo repository.ClientRepo, app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("MAIN PAGE")
	})
	app.Post("/newclient", router.CreateClient)
}

func main() {
	// connect Fiber frameword
	app := fiber.New()

	// connect Postgres database
	db := driver.Connect(host, port, user, password, dbname)
	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

	ClientRepo := implement.NewClientRepo(db.SQL)

	//insert client
	client1 := model.Client{
		ID:     5,
		Name:   "Nha khuyen 5",
		Gender: "Female",
		Email:  "nnn@gm",
	}
	ClientRepo.Insert(client1)

	//select client
	clients, _ := ClientRepo.Select()
	for i := range clients {
		fmt.Println(clients[i])
	}

	setUpRoutes(ClientRepo, app)
	app.Listen(":3000")
}
