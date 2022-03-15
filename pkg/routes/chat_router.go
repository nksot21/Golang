package routes

import (
	"mental-health-api/handler"
	"mental-health-api/pkg/chat"

	"github.com/gofiber/websocket/v2"

	"github.com/gofiber/fiber/v2"
)

func ChatRoutes(app *fiber.App) {
	// MIDDLEWARE
	app.Use("/ws/chat/:userid/:id", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	router := app.Group("/chat")
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("chat data")
	})
	router.Get("/getall/:userid/:id", handler.GetAllMessages)
	router.Get("/:userid/:id", handler.ChatPage)

	runServeWs := chat.ServeWs()
	app.Get("/ws/chat/:userid/:id", runServeWs)
}
