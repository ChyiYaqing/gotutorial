package docker

import (
	"context"

	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetCurrentRunContainers() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalln(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, container := range containers {
		log.Printf("%s\t%s\n", container.ID[:10], container.Image)
	}

}
