package contact

import (
	_ "embed"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/apps/cv/router/navigation"
	"github.com/joshsummers4/golang_cv/apps/cv/router/pages/errors/unexpected"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed get.html
var getHTML string
var getTPL = tpl.Parse("hello page", getHTML)

type getData struct {
	Email    string
	Phone    string
	Address  string
	LinkedIn string
	GitHub   string
}

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

func resolve() (*getData, error) {
	data := &getData{
		Email:    "jswebdev4@gmail.com",
		Address:  "Bristol, UK",
		LinkedIn: "linkedin.com/in/joshua-summers/",
		GitHub:   "github.com/joshsummers4",
	}

	return data, nil
}
