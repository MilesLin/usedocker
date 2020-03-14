package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"os"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	// todo: flags, SSL
	// RestartPolicy
	//   Empty string means not to restart
	//   always Always restart
	//   unless-stopped Restart always except when the user has manually stopped the container
	//   on-failure Restart only when the container exit code is non-zero
	// Authorization, port mapping, container name, imageName

	// todo: enable SSL

	// todo: api port
	imageName := "mileslin/dockerlab"

	// remove image
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image)
		for s, s2 := range image.Labels {
			fmt.Println(s)
			fmt.Println(s2)
		}
		fmt.Println(image.RepoDigests)
		fmt.Println(image.RepoTags)
		break
	}
	//cli.ImageRemove(ctx, )

	// pull image
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
