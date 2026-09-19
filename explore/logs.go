package main

import (
	"context"
	"errors"
	"os"

	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/moby/moby/client"
)

func main() {
	apiClient, err := client.New(
		client.FromEnv,
		client.WithUserAgent("faber/1.0.0"),
	)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	logResult, err := apiClient.ContainerLogs(
		context.Background(),
		"test-nginx",
		client.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow: true,
		},
	)
	if err != nil {
		panic(err)
	}
	defer logResult.Close()

	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, logResult)
	if err != nil && !errors.Is(err, context.Canceled) {
		panic(err)
	}
}
