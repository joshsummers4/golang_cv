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

func SummaryHTML(ctx context.Context) template.HTML {
	return summaryTPL.HTML(ctx, nil)
}
