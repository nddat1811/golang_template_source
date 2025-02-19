package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang_template_source/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	Manager *utils.ConnectionManager
}

type MessageEvent struct {
	Event     string `json:"event"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	ImageData string `json:"image"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebSocketController(manager *utils.ConnectionManager) *WebSocketController {
	return &WebSocketController{Manager: manager}
}

func (ctrl *WebSocketController) HandleWebSocket(c *gin.Context) {
	fmt.Println("WebSocket\n\n\n\n\n\n\n\n\n\nssss")
	roomName := c.Param("room")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()
 
	userID := time.Now().UnixNano() % 10000 // Example user ID generation
	ctrl.Manager.Connect(conn, roomName, int(userID))
	defer ctrl.Manager.Disconnect(conn, roomName)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var event MessageEvent
		if err := json.Unmarshal(msg, &event); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}

		switch event.Event {
		case "typing":
			ctrl.Manager.Broadcast(fmt.Sprintf(`{"event":"typing","user_id":%d}`, event.UserID), roomName, conn)
		case "stop_typing":
			ctrl.Manager.Broadcast(fmt.Sprintf(`{"event":"stop_typing","user_id":%d}`, event.UserID), roomName, conn)
		case "message":
			ctrl.Manager.Broadcast(fmt.Sprintf(`{"event":"message","user_id":%d,"content":"%s"}`, event.UserID, event.Content), roomName, conn)
		default:
			log.Println("Unknown event:", event.Event)
		}
	}
}
