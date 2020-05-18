package interfaces

import (
	"golang-ddd/application"
	"golang-ddd/domain/entity"
	"golang-ddd/utils/helpers"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// MessagesHandler ..
type MessagesHandler struct {
	message application.MessageAppInterface
}

// Message : Define our message object
type Message struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	Channel  string `json:"channel"`
	Message  string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool) // connected clients

var broadcast = make(chan Message) // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// NewMessagesHandler instance from user handler.
func NewMessagesHandler(message application.MessageAppInterface) *MessagesHandler {
	return &MessagesHandler{
		message: message,
	}
}

// ListByChannelName ...
func (handler MessagesHandler) ListByChannelName(c echo.Context) error {
	channel := c.Param("channel")

	messages, err := handler.message.ListByChannelName(channel)

	if err != nil {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Entity Not Found",
		}, http.StatusNotFound, false)
	}

	return helpers.Respond(c, map[string]interface{}{
		"data": messages,
	}, http.StatusOK, true)
}

// HandleConnections ...
func (handler MessagesHandler) HandleConnections(c echo.Context) error {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return err
	}

	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		go func(msg Message, handler MessagesHandler) {
			handler.message.Create(entity.Message{
				Username: msg.Username,
				UserID:   msg.UserID,
				Channel:  msg.Channel,
				Message:  msg.Message,
			})
		}(msg, handler)

		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}

	return nil
}

// HandleWSMessages ...
func HandleWSMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)

			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
