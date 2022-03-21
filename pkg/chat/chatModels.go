package chat

import (
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan Message
	userID string
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}
