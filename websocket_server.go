package main

import "github.com/chyiyaqing/gotutorial.git/pkg/websocket"

func main() {
	go websocket.ServerStart()
	select {}
}
