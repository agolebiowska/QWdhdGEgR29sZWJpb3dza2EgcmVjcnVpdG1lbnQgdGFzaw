package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
)

func Run() {
	conf := config.NewConfig()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.HttpServerPort()),
		Handler: NewRouter(conf),
	}

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Print("Web server shutdown complete.")
		} else {
			log.Fatal(fmt.Printf("The web server shut down unexpectedly: %s", err))
		}
	}

	log.Print(fmt.Printf("Listening on port: %d", conf.HttpServerPort()))
}
