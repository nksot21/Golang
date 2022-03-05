package chat

import (
	"bytes"
	"fmt"
	"log"

	"chatdemo/src/handler"
	"time"

	"github.com/gofiber/websocket/v2"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Message struct {
	SenderID   string
	ReceiverID string
	Content    []byte
}

// READ MESSAGES FROM WEBSOCKET CONNECTION TO HUB
func (c *Client) readPump(conn websocket.Conn, receiverID string) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn = &conn
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		messg := Message{SenderID: c.userID, ReceiverID: receiverID, Content: message}
		id, _ := handler.NewMessage(receiverID, c.userID, message)

		fmt.Println("MessageID: ", id)
		c.hub.broadcast <- messg
	}
}

// PUMPS MESSAGES FROM HUB TO THE WEBSOCKER CONNECTION
func (c *Client) writePump(conn websocket.Conn, receiverID string) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	c.conn = &conn

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}
