package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

type Registrar struct {
	// Interval  time.Duration
	DockerClient *DockerClient
	SRegistry    *ServiceRegistry
}

func (r *Registrar) Init() error {

	cList, err := r.DockerClient.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: filters.NewArgs(
			filters.Arg("ancestor", HelloServiceImageName),
			filters.Arg("status", ContainerRunningState),
		),
	})
	if err != nil {
		return err
	}

	for _, c := range cList {
		r.SRegistry.Add(c.ID, findContainerAddress(c.Ports[0].PublicPort))
	}
	return nil
}

func (r *Registrar) Observe() {
	msgCh, errCh := r.DockerClient.Events(context.Background(), types.EventsOptions{
		Filters: filters.NewArgs(
			filters.Arg("type", "container"),
			filters.Arg("image", HelloServiceImageName),
			filters.Arg("event", "start"),
			filters.Arg("event", "kill"),
		),
	})

	for {
		select {
		case c := <-msgCh:
			fmt.Printf("State of the container %d is %s\n", c.ID, c.Status)
			if c.Status == ContainerKillState {
				r.SRegistry.RemoveByContainerID(c.ID)
			} else if c.Status == ContainerStartState {
				port, err := r.DockerClient.GetContainerPort(context.Background(), c.ID)
				if err != nil {
					fmt.Printf("err getting newly started container port %s\n", err.Error())
					continue
				}
				r.SRegistry.Add(c.ID, findContainerAddress(port))
			}

		case err := <-errCh:
			fmt.Println("Error Docker Event Chan", err.Error())
		}
	}
}

func findContainerAddress(cPort uint16) string {
	return fmt.Sprintf("http://localhost:%d", cPort)
}

/**
func (r *Registrar) Observe() {
	for range time.Tick(r.Interval) {
		cList, _ := r.DockerCLI.ContainerList(context.Background(), types.ContainerListOptions{
			All: true,
		})

		if len(cList) == 0 {
			r.SRegistry.RemoveAll()
			continue
		}

		for _, c := range cList {
			if c.Image != HelloServiceImageName {
				continue
			}

			_, exist := r.SRegistry.GetByContainerID(c.ID)

			if c.State == ContainerRunningState {
				if !exist {
					addr := fmt.Sprintf("http://localhost:%d", c.Ports[0].PublicPort)
					r.SRegistry.Add(c.ID, addr)
				}
			} else {
				if exist {
					r.SRegistry.RemoveByContainerID(c.ID)
				}
			}

		}
	}
}

*/
