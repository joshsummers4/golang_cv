package heading

// heading, and strapline followed by some sort of coding related image (bit like matt farley's cv)
import (
	_ "embed"
	"html/template"
	"context"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed heading.html
var headingHTML string
var headingTPL = tpl.Parse("home heading", headingHTML)

type headingInput struct {
	Header string
	Strapline string
}

func HeadingHTML(ctx context.Context) template.HTML {
	data := &headingInput{
		Header: "Josh Summers - Full Stack Web Developer",
		Strapline: "I code and develop web applications with care and enjoyment",
	}
	return headingTPL.HTML(ctx, data)
}