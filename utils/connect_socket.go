package utils
import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	Conn   *websocket.Conn
	UserID int
}

type ConnectionManager struct {
	activeConnections map[string][]WebSocketConnection
	mutex             sync.Mutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		activeConnections: make(map[string][]WebSocketConnection),
	}
}

func (manager *ConnectionManager) Connect(conn *websocket.Conn, roomName string, userID int) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	if _, exists := manager.activeConnections[roomName]; !exists {
		manager.activeConnections[roomName] = []WebSocketConnection{}
	}

	manager.activeConnections[roomName] = append(manager.activeConnections[roomName], WebSocketConnection{Conn: conn, UserID: userID})
}

func (manager *ConnectionManager) Disconnect(conn *websocket.Conn, roomName string) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	if connections, exists := manager.activeConnections[roomName]; exists {
		newConnections := []WebSocketConnection{}
		for _, connection := range connections {
			if connection.Conn != conn {
				newConnections = append(newConnections, connection)
			}
		}
		manager.activeConnections[roomName] = newConnections

		if len(newConnections) == 0 {
			delete(manager.activeConnections, roomName)
		}
	}
}

func (manager *ConnectionManager) Broadcast(message string, roomName string, excludeConn *websocket.Conn) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	connections := manager.activeConnections[roomName]
	for _, connection := range connections {
		if connection.Conn != excludeConn {
			if err := connection.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				fmt.Printf("Error broadcasting message: %v\n", err)
			}
		}
	}

	// Broadcast to sidebar room if applicable
	if roomName != "sidebar_01" {
		sidebarConnections := manager.activeConnections["sidebar_01"]
		for _, connection := range sidebarConnections {
			if err := connection.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				fmt.Printf("Error broadcasting to sidebar: %v\n", err)
			}
		}
	}
}

func (manager *ConnectionManager) GetConnectedUserIDs(roomName string) []int {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	userIDs := []int{}
	for _, connection := range manager.activeConnections[roomName] {
		userIDs = append(userIDs, connection.UserID)
	}
	return userIDs
}