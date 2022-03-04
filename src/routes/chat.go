package routes

import (
	"chatdemo/src/chat"
	"chatdemo/src/handler"

	"github.com/gofiber/websocket/v2"

	"github.com/gofiber/fiber/v2"
)

func chatPage(c *fiber.Ctx) error {
	return c.SendFile("home.html")
}

func SetUpChatRoutes(app *fiber.App) {
	app.Use("/ws/chat/:userid/:id", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/chat/:userid/:id", chatPage)
	runServeWs := chat.ServeWs()
	app.Get("/:userid/:id", handler.GetAllMessage)
	app.Get("/ws/chat/:userid/:id", runServeWs)
}