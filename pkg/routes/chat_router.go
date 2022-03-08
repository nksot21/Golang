package routes

import (
	"mental-health-api/handler"
	"mental-health-api/pkg/chat"

	"github.com/gofiber/websocket/v2"

	"github.com/gofiber/fiber/v2"
)

func chatPage(c *fiber.Ctx) error {
	//models.GetAllMessages(c.Params("userid"), c.Params("id"))
	return c.SendFile("home.html")
}

func ChatRoutes(app *fiber.App) {
	// MIDDLEWARE
	app.Use("/ws/chat/:userid/:id", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/chat", handler.GetAllMessages)
	app.Get("/chat/:userid/:id", chatPage)
	runServeWs := chat.ServeWs()
	app.Get("/ws/chat/:userid/:id", runServeWs)
}
