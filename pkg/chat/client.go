package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	models "mental-health-api/model"
	"mental-health-api/pkg/const/firestoreCol"
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

type ReceivedMessage struct {
	ReceiverID string
	Content    string
}

type Message struct {
	ID         string
	SenderID   string
	ReceiverID string
	Content    []byte
}

type MessageRes struct {
	ID         string
	SenderID   string
	Sender     UserRef
	ReceiverID string
	Receiver   UserRef
	Content    string
}

type UserRef struct {
	Name      string
	Email     string
	Picture   string
	CreatedAt int64
}

// CREATE MESSAGE: SEND TO CLIENT_WEBSOCKET
func getUserInfo(senderID, receiverID string) (UserRef, UserRef, error) {
	var sender models.User
	var receiver models.User
	var errUser UserRef
	// get users
	if err := sender.GetOne(senderID, ""); err != nil {
		fmt.Println("Get_user_id: ", err)
		return errUser, errUser, err
	}

	if sender.Picture == "" {
		sender.Picture = firestoreCol.DEFAULT_PICTURE
	}

	senderRef := UserRef{
		Name:      sender.Name,
		Email:     sender.Email,
		Picture:   sender.Picture,
		CreatedAt: sender.CreatedAt,
	}
	if err := receiver.GetOne(receiverID, ""); err != nil {
		fmt.Println("Get_user_id: ", err)
		return errUser, errUser, err
	}

	if receiver.Picture == "" {
		receiver.Picture = firestoreCol.DEFAULT_PICTURE
	}

	receiverRef := UserRef{
		Name:      receiver.Name,
		Email:     receiver.Email,
		Picture:   receiver.Picture,
		CreatedAt: receiver.CreatedAt,
	}

	// convert User to []byte
	byteBuffer := new(bytes.Buffer)

	err := json.NewEncoder(byteBuffer).Encode(senderRef)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	senderByte := byteBuffer.Bytes()
	senderByte = bytes.TrimSpace(bytes.Replace([]byte(senderByte), newline, space, -1))

	err = json.NewEncoder(byteBuffer).Encode(receiverRef)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	receiverByte := byteBuffer.Bytes()
	receiverByte = bytes.TrimSpace(bytes.Replace([]byte(receiverByte), newline, space, -1))

	return senderRef, receiverRef, nil
}

// READ MESSAGES FROM WEBSOCKET-CONNECTION TO HUB
func (c *Client) readPump(conn websocket.Conn) {
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

		var receivedMess ReceivedMessage
		if err := json.Unmarshal(message, &receivedMess); err != nil {
			//panic(err)
		}
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		receiverID := receivedMess.ReceiverID

		messageContent := bytes.TrimSpace(bytes.Replace([]byte(receivedMess.Content), newline, space, -1))

		// CREATE A NEW MESSAGE IN DATABASE
		id, _ := models.NewMessage(receiverID, c.userID, []byte(messageContent))

		messg := Message{ID: id, SenderID: c.userID, ReceiverID: receiverID, Content: messageContent}

		fmt.Println("MessageID: ", id)
		c.hub.broadcast <- messg
	}
}

// PUMPS MESSAGES FROM HUB TO THE WEBSOCKER CONNECTION
func (c *Client) writePump(conn websocket.Conn) {
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

			sender, receiver, err := getUserInfo(message.SenderID, message.ReceiverID)
			if err != nil {
				return
			}

			byteBuffer := new(bytes.Buffer)

			messageSended := MessageRes{
				ID:         message.ID,
				SenderID:   message.SenderID,
				Sender:     sender,
				ReceiverID: message.ReceiverID,
				Receiver:   receiver,
				Content:    string(message.Content)}

			err = json.NewEncoder(byteBuffer).Encode(messageSended)
			if err != nil {
				log.Fatal("encode error:", err)
			}

			textByte := byteBuffer.Bytes()
			textByte = bytes.Split(textByte, []byte("\n"))[0]
			w.Write(textByte)

			fmt.Println("mesage: ", string(textByte))

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				messageContent := <-c.send
				w.Write(messageContent.Content)
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
