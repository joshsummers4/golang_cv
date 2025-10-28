package skills

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/design/atoms"
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
	Frontend := []string{"Angular", "TypeScript", "JavaScript", "HTML5", "CSS3", "SASS", "React", "Tailwind CSS", "Storyblok"}
	Backend := []string{"Go", "SQL", "DuckDB", "Parquet files", "MySQL", "DynamoDB", "REST API"}
	DevOps := []string{"Git", "Sourcetree", "Bitbucket", "AWS", "Docker"}
	Testing := []string{"Unit Testing", "Visual Testing", "Agile/Scrum methodolgies"}

	data := skillsData{}
	for _, lang := range Frontend {
		data.Frontend = append(data.Frontend, atoms.PebbleHtml(ctx, lang))
	}
	for _, db := range Backend {
		data.Backend = append(data.Backend, atoms.PebbleHtml(ctx, db))
	}
	for _, tool := range DevOps {
		data.DevOps = append(data.DevOps, atoms.PebbleHtml(ctx, tool))
	}
	for _, test := range Testing {
		data.Testing = append(data.Testing, atoms.PebbleHtml(ctx, test))
	}

	return skillsTPL.HTML(ctx, data)
}
