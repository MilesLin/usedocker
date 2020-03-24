package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// Sample for image name: `stilliard/pure-ftpd:latest`
func RemoveImage(imageNameTag string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

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
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	newStr := buf.String()

	return newStr, nil
}

func PullImageWithAuth(imageName string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	authConfig := types.AuthConfig{
		Username: *cr_acct,
		Password: *cr_pwd,
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return "", err
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	newStr := buf.String()

	return newStr, nil
}
func StopContainer(containerName string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		return "", err
	}
	var isFound = false
	for _, container := range containers {
		for _, name := range container.Names {
			if name == ("/" + containerName) {
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
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})

	if err != nil {
		return "", err
	}

	var isFound = false

	for _, container := range containers {
		for _, name := range container.Names {
			if name == ("/" + containerName) {
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
	portBindings nat.PortMap,
	env []string,
	restartPolicy string,
	bindVolume []Mount) (string, error) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	config := &container.Config{
		Image:        imageName,
		ExposedPorts: exposedPorts,
		Env:          env,
	}
	hostConfig := &container.HostConfig{
		PortBindings: portBindings,
	}

	if len(bindVolume) > 0 {
		for _, m := range bindVolume {
			hostConfig.Mounts = append(hostConfig.Mounts, mount.Mount{
				Type:   mount.Type(m.Type),
				Source: m.Source,
				Target: m.Target,
			})
		}
	}

	// RestartPolicy
	//   Empty string means not to restart
	//   always: Always restart
	//   unless-stopped: Restart always except when the user has manually stopped the container
	//   on-failure: Restart only when the container exit code is non-zero
	if restartPolicy != "" {
		hostConfig.RestartPolicy = container.RestartPolicy{Name: restartPolicy}
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
