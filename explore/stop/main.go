package main

import (
	"context"
	"github.com/moby/moby/client"
)

func main() {
	apiClient, err := client.New(client.FromEnv, client.WithUserAgent("faber/1.0.0"))
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	_, err = apiClient.ContainerStop(context.Background(), "test-nginx", client.ContainerStopOptions{})
	if err != nil {
		panic(err)
	}

	_, err = apiClient.ContainerRemove(context.Background(), "test-nginx", client.ContainerRemoveOptions{})
	if err != nil {
		panic(err)
	}
}