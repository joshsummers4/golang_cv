package router

import (
	"net/http"

	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/contact"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/home"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("GET /", home.GetHandler)
	mux.HandleFunc("GET /contact", contact.GetHandler)
}

func Router() http.Handler {
	return mux
}
