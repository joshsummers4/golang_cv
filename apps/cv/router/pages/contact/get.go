package contact

import (
	_ "embed"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/apps/cv/router/navigation"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/errors/unexpected"
	"github.com/joshsummers4/golang_cv/libs/features/cv"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed get.html
var getHTML string
var getTPL = tpl.Parse("hello page", getHTML)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	data, err := resolve()
	if err != nil {
		unexpected.Handler(w, r, err)
		return
	}

	content := getTPL.HTML(r.Context(), data)

	template := navigation.NavLayout(r, navigation.NavInput{
		Content:     content,
		Title:       "Say Hello",
		Description: "Get in touch with me via email, phone, or social media.",
	})

	io.WriteString(w, template)
}

func resolve() (*cv.ContactInfo, error) {
	data, err := cv.GetContactInfo()
	if err != nil {
		return nil, err
	}
	return data, nil
}
