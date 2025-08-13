package navigation

import (
	_ "embed"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/joshsummers4/golang_cv/libs/ui/page"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed nav.html
var NavHTMl string
var NavTPL = tpl.Parse("navigation layout", NavHTMl)

type NavData struct {
	Content template.HTML
	Year    string
}

type NavInput struct {
	Content     template.HTML
	Title       string
	Description string
}

func NavLayout(r *http.Request, input NavInput) string {
	year := time.Now().Year()
	data := &NavData{
		Content: input.Content,
		Year:    strconv.Itoa(year),
	}
	template := NavTPL.HTML(r.Context(), data)

	pageTemplate := page.Template(&page.PageTemplateInput{
		Title:       input.Title,
		Main:        template,
		Description: input.Description,
	}, r)

	return pageTemplate
}
