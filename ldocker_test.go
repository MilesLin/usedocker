package main

import (
	"github.com/docker/go-connections/nat"
	"testing"
)

func TestUpdateContainer(t *testing.T) {
	StopContainer("golangrun")
	RemoveContainer("golangrun")
	RemoveImage("mileslin/dockerlab:latest")
	PullImage("mileslin/dockerlab:latest")

	portSet := nat.PortSet{"80/tcp": struct{}{}}
	portBindings := nat.PortMap{
		"80/tcp": []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: "8002",
			},
		},
	}
	RunContainer("mileslin/dockerlab:latest",
		"golangrun",
		portSet,
		portBindings,
		[]string{"abc=123", "xyz=999"},
		"always")
}
