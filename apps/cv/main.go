package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joshsummers4/golang_cv/apps/cv/router"
	"github.com/joshsummers4/golang_cv/libs/utils/logger"
)

func main() {
	logger.Init("cv", "v0.01", "Prod", []string{}, logger.LevelTrace)

	mux := router.Router()
	//add logging at app level for dev mode
	muxLogging := logger.LoggingHandler(
		[]string{"/public", "/favicon.ico", "/.well-known/appspecific/com.chrome.devtools.json"},
	)(mux)
	port := "8080"
	logger.Info(context.Background(), "server starting", []string{"server"}, nil)

	logger.Debug(context.Background(), fmt.Sprintf("server listening at http://localhost:%s", port), []string{"server"}, nil, nil)

	if err := http.ListenAndServe(":"+port, muxLogging); err != nil {
		logger.Error(context.Background(), "server has failed", err, []string{"server"}, nil)
	}

}
