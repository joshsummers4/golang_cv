package skills

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/design/atoms"
	"github.com/joshsummers4/golang_cv/libs/features/cv"
	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed skills.html
var skillsHTML string
var skillsTPL = tpl.Parse("skills section", skillsHTML)

type skillsData struct {
	Frontend []template.HTML
	Backend  []template.HTML
	DevOps   []template.HTML
	Testing  []template.HTML
}

func SkillsHTML(ctx context.Context) template.HTML {
	skills, _ := cv.GetSkills("")

	data := skillsData{}
	for _, lang := range skills.Frontend {
		data.Frontend = append(data.Frontend, atoms.PebbleHtml(ctx, lang))
	}
	for _, db := range skills.Backend {
		data.Backend = append(data.Backend, atoms.PebbleHtml(ctx, db))
	}
	for _, tool := range skills.DevOps {
		data.DevOps = append(data.DevOps, atoms.PebbleHtml(ctx, tool))
	}
	for _, test := range skills.Testing {
		data.Testing = append(data.Testing, atoms.PebbleHtml(ctx, test))
	}

	return skillsTPL.HTML(ctx, data)
}
