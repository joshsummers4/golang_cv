package notfound

import (
	_ "embed"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/apps/cv/router/navigation"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed notfound.html
var notfoundHTML string
var notfoundTPL = tpl.Parse("notfound error", notfoundHTML)

func Handler(w http.ResponseWriter, r *http.Request) {
	templ := notfoundTPL.HTML(r.Context(), nil)

	input := navigation.NavInput{
		Content: templ,
		Title:   "Unexpected error",
	}
	navTemplate := navigation.NavLayout(r, input)

	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, navTemplate)
}
