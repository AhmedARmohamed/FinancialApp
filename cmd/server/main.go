package main

import (
	"github.com/AhmedARmohamed/FINANCIAL-APP-BACKEND/internal/api"
	"github.com/AhmedARmohamed/FINANCIAL-APP-BACKEND/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

//Create Server object and start listener
func main() {

	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithField("version", config.Version).Debug("Starting server")

	router, err := api.NewRouter()
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}
	const addr = "0.0.0.0:8080"
	server := http.Server{
		Handler: router,
		Addr: addr,
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.WithError(err).Error("Server failed.")
	}
}
