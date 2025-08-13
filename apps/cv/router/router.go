package router

import (
	"net/http"

	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/contact"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/home"
	"github.com/joshsummers4/golang_cv/libs/utils/logger"
)

var mux *http.ServeMux = http.NewServeMux()

func Router() http.Handler {

	mux.HandleFunc("GET /", home.GetHandler)
	mux.HandleFunc("GET /contact", contact.GetHandler)

	loggingMux := logger.LoggingHandler([]string{"/public", "/favicon.ico", "/.well-known/appspecific/com.chrome.devtools.json"})(mux)

	return loggingMux
}
