package main

import (
	"context"
	"fmt"
	"log"

	"github.com/moby/moby/client"
)


func main() {
	apiClient, err := client.New(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	result, err := apiClient.ContainerList(context.Background(), client.ContainerListOptions{ All: true })
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %-22s %s\n", "ID", "STATUS", "IMAGE")
	for _, ctr := range result.Items {
		fmt.Printf("%s %-22s %s\n", ctr.ID, ctr.Status, ctr.Image)
	}
}