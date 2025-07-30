package handler

import (
    "net/http"

	cvrouter "github.com/joshsummers4/golang_cv/apps/cv/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	//for running in production using vercel
    cvrouter.Router().ServeHTTP(w, r)
}