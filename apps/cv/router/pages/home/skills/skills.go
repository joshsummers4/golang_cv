package skills

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
	"github.com/joshsummers4/golang_cv/libs/design/atoms"
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
	Languages:=    []string{"Go", "TypeScript", "JavaScript", "Angular", "AlpineJS", "HTML5", "CSS3", "SASS"}
	Databases:=    []string{"SQL", "DuckDB", "Parquet files", "MySQL", "Microsoft Access", "DynamoDB"}
	Tools:=        []string{"Git", "Sourcetree", "Bitbucket", "AWS", "Storyblok", "Tailwind CSS"}
	Testing:=      []string{"Unit Testing", "Visual Testing", "Agile/Scrum methodolgies"}

	data:= skillsData{}
	for _, lang := range Languages {
		data.Languages = append(data.Languages, atoms.PebbleHtml(ctx, lang))
	}
	for _, db := range Databases {
		data.Databases = append(data.Databases, atoms.PebbleHtml(ctx, db))
	}
	for _, tool := range Tools {
		data.Tools = append(data.Tools, atoms.PebbleHtml(ctx, tool))
	}
	for _, test := range Testing {
		data.Testing = append(data.Testing, atoms.PebbleHtml(ctx, test))
	}
	
	return skillsTPL.HTML(ctx, data)
}
