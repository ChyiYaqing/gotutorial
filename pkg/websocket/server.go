package websocket

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func ServerStart() {
	http.HandleFunc("/ws", handleWebSocket)
	http.ListenAndServe("127.0.0.1:8888", nil)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Server=> Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	pingInterval := time.Second * 5
	pongWait := time.Second * 10

	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("Error sending message:", err)
				break
			}
		default:
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Println("Server=> Connection closed unexpectedly:", err)
				} else {
					log.Println("Server=> Error reading message:", err)
				}
				return
			}
			switch messageType {
			case websocket.TextMessage:
				log.Printf("Server=> Received text message: %s\n", string(message))
				// 回复消息
				if err := conn.WriteMessage(messageType, message); err != nil {
					log.Println("Server=> Error sending text message:", err)
					break
				}
			case websocket.BinaryMessage:
				log.Printf("Server=> Received binary message: %s\n", string(message))
				// 回复消息
				if err := conn.WriteMessage(messageType, message); err != nil {
					log.Println("Server=> Error sending binary message:", err)
					break
				}
			case websocket.PingMessage:
				// 获取Ping消息的Reader对象
				log.Println("Server=> Received PingMessage")
				if err := conn.WriteMessage(websocket.PongMessage, []byte{}); err != nil {
					log.Println("Server=> Error sending PongMessage message:", err)
					break
				}
			case websocket.PongMessage:
				log.Println("Server=> Received CloseMessage")
				conn.SetReadDeadline(time.Now().Add(pongWait))
			case websocket.CloseMessage:
				log.Println("Server=> Closing connection...")
				return
			}
		}
	}
}
