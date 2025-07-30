package skills

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/utils/tpl"
)

//go:embed skills.html
var skillsHTML string
var skillsTPL = tpl.Parse("skills section", skillsHTML)

type skillsData struct {
	Languages    []string
	Databases    []string
	Tools        []string
	Integrations []string
	Testing      []string
}

func SkillsHTML(ctx context.Context) template.HTML {
	data := &skillsData{
		Languages:    []string{"Go", "JavaScript", "TypeScript", "Angular", "HTML5", "CSS3", "Sass", "PHP", "C#", "Python"},
		Databases:    []string{"SQL", "DuckDB", "Parquet files", "MySQL", "Microsoft Access", "DynamoDB"},
		Tools:        []string{"Git", "Github", "Sourcetree", "Bitbucket", "AWS (app runner, cognito, AVP, s3, cloudwatch)", "Storybook", "Tailwind CSS", "Vercel"},
		Integrations: []string{"Storyblok", "WordPress", "Paddle billing webhooks", "ZOHO One", "UiPath"},
		Testing:      []string{"Unit Testing", "Visual Testing", "Agile/Scrum methodolgies"},
	}
	return skillsTPL.HTML(ctx, data)
}
