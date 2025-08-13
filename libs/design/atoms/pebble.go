package atoms

import (
	_ "embed"
	"html/template"
	"context"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed pebble.html
var pebbleHTML string
var pebbleTPL = tpl.Parse("pebble atom", pebbleHTML)

func PebbleHtml(ctx context.Context, text string) template.HTML {
	return pebbleTPL.HTML(ctx, text)
}