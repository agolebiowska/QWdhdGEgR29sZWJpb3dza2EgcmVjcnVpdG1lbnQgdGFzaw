package server

import (
	"fmt"
	"net/http"

	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
	"github.com/sirupsen/logrus"
)

func Run() {
	conf := config.NewConfig()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.HttpServerPort()),
		Handler: NewRouter(conf),
	}

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logrus.Print("Web server shutdown complete.")
		} else {
			logrus.Fatal(fmt.Printf("The web server shut down unexpectedly: %s", err))
		}
	}

	logrus.Print(fmt.Printf("Listening on port: %d", conf.HttpServerPort()))
}
