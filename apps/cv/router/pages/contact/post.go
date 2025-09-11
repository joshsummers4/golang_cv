package contact

import (
	_ "embed"
	"io"
	"net/http"

	"github.com/joshsummers4/golang_cv/libs/features/cv"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed post.html
var PostHTML string
var PostTPL = tpl.Parse("contact form response", PostHTML)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Handle POST request for contact form submission
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// Here you would typically process the form data, e.g., send an email or store it in a database.
	// store it in a database.
	err := cv.AddContact(r.Context(), name, email, message)
	if err != nil {
		http.Error(w, "Failed to process contact form", http.StatusInternalServerError)
		return
	}

	// replace form with a thank you message
	// Send a response back to the client
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, PostTPL.String(r.Context(), nil))
}
