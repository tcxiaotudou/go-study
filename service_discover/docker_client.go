package main

import (
	"context"
	"github.com/docker/docker/client"
)

const (
	HelloServiceImageName = "hello"
	ContainerRunningState = "running"
	ContainerKillState    = "kill"
	ContainerStartState   = "start"
)

type DockerClient struct {
	*client.Client
}

func NewDockerClient() (*DockerClient, error) {
	dockerCLI, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	return &DockerClient{dockerCLI}, nil
}

func (dc DockerClient) GetContainerPort(ctx context.Context, id string) (uint16, error) {

	return 0, nil
}
