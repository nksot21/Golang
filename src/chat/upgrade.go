package chat

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// THEM USER ID

func ServeWs() func(*fiber.Ctx) error {
	// tạo 1 connect socket với chỉ số được khai báo
	var ConnectWS = websocket.New(func(c *websocket.Conn) {
		var wg sync.WaitGroup
		wg.Add(100)
		log.Println(c.Locals("allowed"))

		receiverID := c.Params("id")
		client := &Client{hub: HubConn, conn: c, send: make(chan []byte, 256), userID: c.Params("userid")}
		client.hub.register <- client

		//var connCh = make(chan *websocket.Conn, 2)

		fmt.Println("New client: ", c)
		var conn = *c
		//connCh <- c
		//connCh <- c

		go client.readPump(conn, receiverID)
		go client.writePump(conn, receiverID)
		wg.Wait()
		fmt.Print("End!!!!")
	},
		websocket.Config{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	)

	return ConnectWS
}
