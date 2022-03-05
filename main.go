package main

import (
	"chatdemo/src/chat"
	"chatdemo/src/firebase"

	"chatdemo/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	//CONNECT DATABASE
	firebase.ConnectFirebase()
	firebase.ConnectFirestore()

	//CREATE CHAT HUB
	go chat.HubConn.Run()
	routes.SetUpChatRoutes(app)
	routes.SetUpUserRoutes(app)

	app.Listen(":3000")
}
