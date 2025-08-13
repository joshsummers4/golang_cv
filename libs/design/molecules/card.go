package molecules

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed card.html
var cardHTML string
var cardTPL = tpl.Parse("card molecule", cardHTML)

type CardData struct {
	Title string
	Description string
	Languages []string
}

func Card(ctx context.Context, data CardData) template.HTML {
	return cardTPL.HTML(ctx, data)
}