package experience

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/design/atoms"
	"github.com/joshsummers4/golang_cv/libs/design/organisms"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed experience.html
var experienceHTML string
var experienceTPL = tpl.Parse("experience section", experienceHTML)

type ExperienceData struct {
	Text      string
	Languages []template.HTML
}

func Experience(ctx context.Context) template.HTML {
	return experienceTPL.HTML(ctx, nil)
}

func ExperienceHTML(ctx context.Context) template.HTML {
	var roles []organisms.FocusCard

	altPebbles := []template.HTML{
		atoms.PebbleHtml(ctx, "Go"),
		atoms.PebbleHtml(ctx, "Angular.js"),
		atoms.PebbleHtml(ctx, "Javascript"),
		atoms.PebbleHtml(ctx, "Typescript"),
		atoms.PebbleHtml(ctx, "HTML"),
		atoms.PebbleHtml(ctx, "CSS"),
		atoms.PebbleHtml(ctx, "Sass"),
		atoms.PebbleHtml(ctx, "SQL"),
	}
	altData := ExperienceData{
		Text:      "",
		Languages: altPebbles,
	}

	alterian := organisms.FocusCard{
		Title:       "Web Developer",
		Description: "Alterian (2022 - Present)",
		Template:    experienceTPL.HTML(ctx, altData),
	}

	alterian2 := organisms.FocusCard{
		Title:       "UI",
		Description: "Alterian (2022 - Present)",
		Template:    experienceTPL.HTML(ctx, altData),
	}
	alterian3 := organisms.FocusCard{
		Title:       "App",
		Description: "Alterian (2022 - Present)",
		Template:    experienceTPL.HTML(ctx, altData),
	}

	roles = append(roles, alterian)
	roles = append(roles, alterian2)
	roles = append(roles, alterian3)

	return organisms.FocusCardsHTML(ctx, roles)
}
