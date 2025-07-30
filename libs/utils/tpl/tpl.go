package tpl

import (
	"bytes"
	"context"
	"html/template"

	"github.com/joshsummers4/golang_cv/libs/utils/logger"
)

type Template struct {
	*template.Template
}

func Parse(name string, html string) *Template {
	return &Template{
		Template: template.Must(
			template.New(name).Parse(html),
		),
	}
}

func ParseWithFuncs(name string, html string, funcs template.FuncMap) *Template {
	return &Template{
		Template: template.Must(
			template.New(name).Funcs(funcs).Parse(html),
		),
	}
}

func String(ctx context.Context, t *Template, body any) (string, error) {
	var output bytes.Buffer

	err := t.Execute(&output, body)

	if err != nil {
		// fmt.Println("error parsing template:", err)
		logger.Error(ctx, "error parsing template", err, []string{"server"}, nil)
		return "", err
	}

	return output.String(), nil
}

func HTML(ctx context.Context, tpl *Template, body any) (template.HTML, error) {
	output, err := String(ctx, tpl, body)

	if err != nil {
		// fmt.Println("error parsing template:", err)
		logger.Error(ctx, "error parsing template", err, []string{"server"}, nil)
		return "", err
	}

	return template.HTML(output), err
}

func Must[R string | template.HTML](t R, err error) R {
	return t
}

func PreRender(name string, html string, content any) template.HTML {
	tpl := Parse(name, html)

	return Must(HTML(context.Background(), tpl, content))
}

func (t *Template) HTML(ctx context.Context, body any) template.HTML {
	return Must(HTML(ctx, t, body))
}

func (t *Template) String(ctx context.Context, body any) string {
	return Must(String(ctx, t, body))
}
