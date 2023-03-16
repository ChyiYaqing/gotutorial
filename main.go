package main

import (
	"github.com/chyiyaqing/gotutorial.git/pkg/docker"
	"github.com/chyiyaqing/gotutorial.git/pkg/mutex"
)

func main() {
	mutex.AdditionCount()

	// For example, to list running containers (the equivalent of docker ps)
	docker.GetCurrentRunContainers()
}
