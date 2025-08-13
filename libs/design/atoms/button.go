package atoms

import (
	"context"
	_ "embed"
	"html/template"
	"strings"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed button.html
var buttonHTML string
var buttonTPL = tpl.Parse("button atom", buttonHTML)

type Button struct {
  ID string
  Label string
  Type string // filled, outlined, text
  Link string // URL to bavigatet to 
}

func ButtonHtml(ctx context.Context, button Button) template.HTML {
  return buttonTPL.HTML(ctx, button)
}

func (b *Button) Classes() string {
	//include tailwind classes
  classes := []string{"button"}

  return strings.Join(classes, " ")
}