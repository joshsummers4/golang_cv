package summary

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed summary.html
var summaryHTML string
var summaryTPL = tpl.Parse("summary section", summaryHTML)

type SummaryData struct {
	Email    string
	Phone    string
	Address  string
	LinkedIn string
	GitHub   string
}

func SummaryHTML(ctx context.Context) template.HTML {
	data := &SummaryData{
		Email:    "scrim.smogs_5x@icloud.com",
		Address:  "Bristol, UK",
		LinkedIn: "https://www.linkedin.com/in/joshua-summers/",
		GitHub:   "https://github.com/joshsummers4",
	}
	return summaryTPL.HTML(ctx, data)
}
