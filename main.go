package main

import (
	"chatdemo/src/chat"
	"chatdemo/src/firebase"

	"chatdemo/src/routes"

	"github.com/gofiber/fiber/v2"
)

func runChat() {
	hub := chat.NewHub()

	go hub.Run()
}

func main() {

	app := fiber.New()

	firebase.ConnectFirebase()
	firebase.ConnectFirestore()

	//CREATE CHAT HUB
	go chat.HubConn.Run()
	routes.SetUpChatRoutes(app)
	routes.SetUpUserRoutes(app)

	app.Listen(":3000")
}
