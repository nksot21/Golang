package chat

var HubConn = NewHub()

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case messg := <-h.broadcast:
			for client := range h.clients {
				if client.userID == messg.ReceiverID || client.userID == messg.SenderID {
					select {
					case client.send <- messg:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
