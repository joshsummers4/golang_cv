package notfound

import (
	_ "embed"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/libs/ui/page"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed notfound.html
var notfoundHTML string
var notfoundTPL = tpl.Parse("notfound error", notfoundHTML)

func Handler(w http.ResponseWriter, r *http.Request) {
	templ := notfoundTPL.HTML(r.Context(), nil)

	page := page.Template(&page.PageTemplateInput{
		Title: "Page not found",
		Main:  templ,
	}, r)

	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, page)
}
