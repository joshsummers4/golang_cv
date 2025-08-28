package heading

// heading, and strapline followed by some sort of coding related image (bit like matt farley's cv)
import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed heading.html
var headingHTML string
var headingTPL = tpl.Parse("home heading", headingHTML)

type headingInput struct {
	Header    string
	Strapline string
	Email     string
	LinkedIn  string
	GitHub    string
}

func HeadingHTML(ctx context.Context) template.HTML {
	data := &headingInput{
		Header:    "Josh Summers",
		Strapline: "Web Developer | Full Stack Developer | Frontend Developer",
		Email:     "scrim.smogs_5x@icloud.com",
		LinkedIn:  "linkedin.com/in/joshua-summers/",
		GitHub:    "github.com/joshsummers4",
	}
	return headingTPL.HTML(ctx, data)
}
