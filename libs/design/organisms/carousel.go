package organisms

import (
	_ "embed"
	"html/template"
	"encoding/json"
	"context"

	"github.com/joshsummers4/golang_cv/libs/utils/logger"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed carousel.html
var carouselHTML string
var carouselTPL = tpl.Parse("carousel organism", carouselHTML)

type CarouselData struct {
	ItemsJson string
}

func CarouselsHTML(ctx context.Context, items []template.HTML) template.HTML {
	data := &CarouselData{}

	itemsJson, err := json.Marshal(items)
	if err != nil {
		logger.Error(ctx, "Carousel: error marshalling cards", err, []string{"organisms", "Carousel"})
		return ""
	}

	data.ItemsJson = string(itemsJson)

	return carouselTPL.HTML(ctx, data)
}