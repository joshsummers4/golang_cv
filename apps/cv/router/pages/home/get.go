package home

import (
	_ "embed"
	"html/template"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/errors/notfound"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/errors/unexpected"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/home/contact"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/home/skills"
	"github.com/joshsummers4/golang_cv/libs/ui/page"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed get.html
var getHTML string
var getTPL = tpl.Parse("home page", getHTML)

type getInput struct {
	Sections []template.HTML
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notfound.Handler(w, r)
		return
	}
	data, err := resolve(r)
	if err != nil {
		unexpected.Handler(w, r, err)
		return
	}

	content := getTPL.HTML(r.Context(), data)

	pageTemplate := page.Template(&page.PageTemplateInput{
		Title: "Home",
		Main:  content,
	}, r)

	io.WriteString(w, pageTemplate)
}

func resolve(r *http.Request) (*getInput, error) {
	//get sections for home page
	data := &getInput{}

	contacts := contact.ContactHTML(r.Context())
	data.Sections = append(data.Sections, contacts)

	skill := skills.SkillsHTML(r.Context())
	data.Sections = append(data.Sections, skill)
	return data, nil
}
