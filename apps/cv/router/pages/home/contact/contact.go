package contact

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed contact.html
var contactHTML string
var contactTPL = tpl.Parse("contact section", contactHTML)

type contactData struct {
	Email    string
	Phone    string
	Address  string
	LinkedIn string
	GitHub   string
}

func ContactHTML(ctx context.Context) template.HTML {
	data := &contactData{
		Email:    "scrim.smogs_5x@icloud.com",
		Address:  "Bristol, UK",
		LinkedIn: "https://www.linkedin.com/in/joshua-summers/",
		GitHub:   "https://github.com/joshsummers4",
	}
	return contactTPL.HTML(ctx, data)
}
