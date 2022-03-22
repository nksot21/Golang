package chat

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ServeWs() func(*fiber.Ctx) error {
	var ConnectWS = websocket.New(func(c *websocket.Conn) {
		var wg sync.WaitGroup
		wg.Add(100)
		log.Println(c.Locals("allowed"))
		//receiverID := c.Params("id")
		client := &Client{hub: HubConn, conn: c, send: make(chan Message, 256), userID: c.Params("userid")}
		client.hub.register <- client

		fmt.Println("New client")
		var conn = *c

		go client.readPump(conn)
		go client.writePump(conn)
		wg.Wait()
	},
		websocket.Config{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	)
	return ConnectWS
}
