package main

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := "mileslin/dockerlab"
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	config := &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"80/tcp": struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8001",
				},
			},
		},
	}
	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, "golangrun")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

//
//func main22() {
//	r := gin.Default()
//
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": "pong",
//		})
//	})
//
//	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}
