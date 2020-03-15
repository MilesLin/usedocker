package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

var (
	ctx context.Context
	cli *client.Client
)

func init() {
	var err error
	ctx = context.Background()
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
}

// Sample for image name: `stilliard/pure-ftpd:latest`
func RemoveImage(imageNameTag string) (string, error) {
	images, err := cli.ImageList(ctx, types.ImageListOptions{})

	if err != nil {
		return "", err
	}

	var isfound = false
	var message = ""
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == imageNameTag {
				isfound = true
				res, err := cli.ImageRemove(ctx, image.ID, types.ImageRemoveOptions{
					Force:         true,
					PruneChildren: true,
				})

				if err != nil {
					return "", err
				}

				for _, re := range res {
					message += fmt.Sprintf("Image %s deleted\n", re.Deleted)
				}

				break
			}
		}
	}

	if !isfound {
		return "", errors.New(fmt.Sprintf("Image %s not found", imageNameTag))
	} else {
		return message, nil
	}

}

// Sample for image name: `stilliard/pure-ftpd:latest`
func PullImage(imageName string) (string, error) {
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	newStr := buf.String()

	return newStr, nil
}

func StopContainer(containerName string) (string, error) {

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		return "", err
	}
	var isFound = false
	for _, container := range containers {
		for _, name := range container.Names {
			if name == containerName {
				isFound = true
				if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
					return "", err
				}
			}
		}
	}

	if isFound {
		return "Stop Container Success", nil
	} else {
		return "", errors.New("container not found")
	}

}

func RemoveContainer(containerName string) (string, error) {

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})

	if err != nil {
		return "", err
	}

	var isFound = false

	for _, container := range containers {
		for _, name := range container.Names {
			if name == containerName {
				isFound = true
				fmt.Print("Removing container ", container.ID[:10], "... ")
				if err := cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{Force: true}); err != nil {
					return "", err
				}
			}
		}
	}

	if isFound {
		return "Removed container", nil
	} else {
		return "", errors.New("container not found")
	}

}

// Example of argument
// imageName: mileslin/dockerlab:latest
// containerName: app
// exposedPorts: nat.PortSet{"80/tcp": struct{}{}}
// portBindings:
//      nat.PortMap{
//      	"80/tcp": []nat.PortBinding{
//      		{
//      			HostIP:   "0.0.0.0",
//      			HostPort: "8001",
//      		},
//      	},
//      }
func RunContainer(
	imageName,
	containerName string,
	exposedPorts nat.PortSet,
	portBindings nat.PortMap) (string, error) {

	config := &container.Config{
		Image:        imageName,
		ExposedPorts: exposedPorts,
	}
	hostConfig := &container.HostConfig{
		PortBindings:  portBindings,
		RestartPolicy: container.RestartPolicy{Name: "always"},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	return fmt.Sprintf("Container %s is running\n", containerName), nil
}
