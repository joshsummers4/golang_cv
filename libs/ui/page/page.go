package page

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed page.html
var pageHTML string
var pageTpl = tpl.ParseWithFuncs("pagetemplate", pageHTML, template.FuncMap{
	"iconsHref": func(icons []string) template.HTMLAttr {
		googleFontsHref := "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"
		iconNames := strings.Join(icons, ",")

		return template.HTMLAttr(fmt.Sprintf("%s&icon_names=%s&swap=true", googleFontsHref, iconNames))
	},
})



type PageTemplateInput struct {
	AppName    string
	Title      string
	Main       template.HTML
	Loading    template.HTML
}

func Template(input *PageTemplateInput, r *http.Request) string {
	//input.Loading = loading.Atom() TODO
	input.AppName = "Josh Summers' Golang CV"
	output := pageTpl.String(r.Context(), input)

	return output
}