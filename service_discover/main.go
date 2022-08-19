package main

import (
	"log"
	"net/http"
)

func main() {

	register := ServiceRegistry{}
	register.Init()

	dockerClient, err := NewDockerClient()
	if err != nil {
		panic(err)
	}

	registrat := Registrar{
		DockerClient: dockerClient,
		SRegistry:    &register,
	}
	if err = registrat.Init(); err != nil {
		panic(err)
	}

	go registrat.Observe()

	app := Application{Sregistry: &register}
	http.HandleFunc("/reverse-proxy", app.Handle)

	log.Fatalln(http.ListenAndServe(":3000", nil))
}
