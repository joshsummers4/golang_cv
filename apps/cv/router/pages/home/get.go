package home

import (
	_ "embed"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/libs/ui/page"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed get.html
var getHTML string
var getTPL = tpl.Parse("home page", getHTML)

type getInput struct {
	Title string
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	content := getTPL.HTML(r.Context(), getInput{})

	input := &page.PageTemplateInput{
		Title: "Home",
		Main:  content,
	}

	output := page.Template(input, r)

	io.WriteString(w, output)
}
