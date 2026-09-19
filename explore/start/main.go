package main

import (
	"context"
	"github.com/moby/moby/client"
	"github.com/moby/moby/api/types/container"
)

func main() {
	apiClient, err := client.New(client.FromEnv, client.WithUserAgent("faber/1.0.0"))
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	result, err := apiClient.ImagePull(context.Background(), "nginx:alpine", client.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	result.Wait(context.Background())

	createResult, err := apiClient.ContainerCreate(
		context.Background(), 
		client.ContainerCreateOptions{
			Name: "test-nginx",
			Config: &container.Config{
				Image: "nginx:alpine",
			},
		},
	)
	if err != nil {
		panic(err)
	}

	_, err = apiClient.ContainerStart(context.Background(), createResult.ID, client.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
}