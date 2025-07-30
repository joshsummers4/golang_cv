package unexpected

import (
	_ "embed"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/libs/ui/page"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

type unexpectedData struct {
	Error error
}

//go:embed unexpected.html
var unexpectedHTML string
var unexpectedTPL = tpl.Parse("unexpected error", unexpectedHTML)

// Handler handles unexpected errors by rendering a template with the error message.
func Handler(w http.ResponseWriter, r *http.Request, err error) {
	ctx := r.Context()

	data := &unexpectedData{
		Error: err,
	}

	templ := unexpectedTPL.HTML(ctx, data)

	page := page.Template(&page.PageTemplateInput{
		Title: "Unexpected Error",
		Main:  templ,
	}, r)

	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, page)
}
