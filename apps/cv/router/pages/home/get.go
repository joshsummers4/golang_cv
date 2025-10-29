package home

import (
	_ "embed"
	"html/template"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/apps/cv/router/navigation"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/errors/notfound"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/errors/unexpected"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/home/experience"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/home/heading"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/home/skills"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/home/summary"
	"github.com/joshsummers4/golang_cv/libs/features/cv"
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

	input := navigation.NavInput{
		Content:     content,
		Title:       "Josh Summers Golang CV",
		Description: "Josh Summers Golang CV. Here you can find my skills, experience, contact info, and more.",
	}
	navTemplate := navigation.NavLayout(r, input)

	io.WriteString(w, navTemplate)
}

func resolve(r *http.Request) (*getInput, error) {
	data := &getInput{}

	contactinfo, err := cv.GetContactInfo()
	if err != nil {
		return nil, err
	}

	header := heading.HeadingHTML(r.Context(), contactinfo)
	summary := summary.SummaryHTML(r.Context())
	skills := skills.SkillsHTML(r.Context())
	experience := experience.Experience(r.Context())
	data.Sections = []template.HTML{header, summary, skills, experience}

	return data, nil
}
