package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
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

	// See https://blog.bracebin.com/achieving-perfect-ssl-labs-score-with-go
	tlsConfig := func(srv *http.Server) {
		srv.TLSConfig = &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}
	}

	// listen and serve
	// webServer := server.CreateServer(restHandler.GetRouter(), ":"+os.Getenv("HTTP_PORT"))
	webServer := infra.CreateServer(restHandler.GetRouter(), ":3000", tlsConfig)
	log.Println("starting server...")
	return webServer.ListenAndServe()
}
