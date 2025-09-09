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
	Languages []template.HTML
	Databases []template.HTML
	Tools     []template.HTML
	Testing   []template.HTML
}

func SkillsHTML(ctx context.Context) template.HTML {
	skills := cv.GetSkills(ctx)
	data := skillsData{}
	for _, s := range skills {
		switch s.Type {
		case "Language":
			data.Languages = append(data.Languages, atoms.PebbleHtml(ctx, s.Skill))
		case "Database":
			data.Databases = append(data.Databases, atoms.PebbleHtml(ctx, s.Skill))
		case "Tools":
			data.Tools = append(data.Tools, atoms.PebbleHtml(ctx, s.Skill))
		case "Testing":
			data.Testing = append(data.Testing, atoms.PebbleHtml(ctx, s.Skill))
		}
	}
	return skillsTPL.HTML(ctx, &data)
}
