package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marvincaspar/go-web-app-boilerplate/pkg/http/rest"
	"github.com/marvincaspar/go-web-app-boilerplate/pkg/infra"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {
	// init logger
	// level 0 is a info level ,so debug level doesn't show
	// debug level can show in level -1
	logger, err := infra.CreateLogger(0)
	if err != nil {
		return err
	}

	// init storage

	// init repositories with given logger and storage

	// init services with given logger and repository

	// setup routes
	restHandler := rest.CreateHandler(logger)
	restHandler.NewHealthCheckHandler()

	// listen and serve
	// webServer := server.CreateServer(restHandler.GetRouter(), ":"+os.Getenv("HTTP_PORT"))
	webServer := infra.CreateServer(restHandler.GetRouter(), ":3000")
	log.Println("starting server...")
	return webServer.ListenAndServe()
}
