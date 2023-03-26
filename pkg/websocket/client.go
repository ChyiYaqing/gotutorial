package websocket

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

func ClientStart() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8888/ws", nil)
	if err != nil {
		fmt.Println("Failed to connect to WebSocket sever:", err)
		return
	}
	defer conn.Close()

	pingInterval := time.Second * 5 // 心跳间隔为5秒
	pongWait := time.Second * 10    // 等待服务端回复的超时时间为10秒

	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("Client=> Error sending ping message:", err)
				return
			}
			log.Println("Client=> Success sending ping message")
		case <-interrupt:
			fmt.Println("Client=> Interrupt signal received, closing connection...")
			if err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				log.Println("Client=> Error sending close message:", err)
				return
			}
			return
		default:
			// 设置读取超时时间
			conn.SetReadDeadline(time.Now().Add(pongWait))
			messageType, message, err := conn.ReadMessage()
			// 客户端代码中捕获异常并正常关闭连接
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Println("Client=> Connection closed unexpectedly:", err)
				} else {
					log.Println("Client=> Error reading message:", err)
				}
				break
			}
			switch messageType {
			case websocket.TextMessage:
				log.Printf("Client=> Received Text message: %s\n", string(message))
			case websocket.BinaryMessage:
				log.Printf("Client=> Received Binary message: %s\n", string(message))
			case websocket.PingMessage:
				log.Printf("Client=> Received Ping message: %s\n", string(message))
				// 回复Pong消息
				if err := conn.WriteMessage(websocket.PongMessage, []byte{}); err != nil {
					log.Println("Client=> Error sending pong message:", err)
					return
				}
			case websocket.PongMessage:
				// 忽略PongMessage
				log.Println("Client=> Received Pong message")
				// 重置读取超时时间
				conn.SetReadDeadline(time.Now().Add(pongWait))
			case websocket.CloseMessage:
				log.Println("Received CloseMessage")
				return
			}
		}
	}
}
