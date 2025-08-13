package organisms

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed focuscards.html
var focusCardsHTML string
var focusCardsTPL = tpl.Parse("focuscards", focusCardsHTML)

type FocusCard struct {
	Title       string
	Description string
	Template    template.HTML
}

type FocusCardsData struct {
	Items []FocusCard
	Focus string
}

// FocusCardsHTML returns the HTML for the focus cards section.
func FocusCardsHTML(ctx context.Context, cards []FocusCard) template.HTML {
	data := &FocusCardsData{
		Items: cards,
		Focus: cards[0].Title,
	}

	return focusCardsTPL.HTML(ctx, data)
}
