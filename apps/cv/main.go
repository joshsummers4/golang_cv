package main

import (
	"fmt"
	"net/http"

	"github.com/joshsummers4/golang_cv/apps/cv/router"
	"github.com/joshsummers4/golang_cv/config"
)

func main() {
	mux := router.Router()

	config.GetEnvironmentVariables()

	muxLogging := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		mux.ServeHTTP(w, r)
	})
	port := "8080"
	fmt.Printf("server starting\n")

	fmt.Printf("server listening at http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, muxLogging); err != nil {
		fmt.Printf("server has failed: %v\n", err)
	}

}
