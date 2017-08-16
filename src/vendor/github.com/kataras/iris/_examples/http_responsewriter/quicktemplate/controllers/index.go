package controllers

import (
	"github.com/kataras/iris/_examples/http_responsewriter/quicktemplate/templates"

	"github.com/kataras/iris/context"
)

// Index renders our ../templates/index.qtpl file using the compiled ../templates/index.qtpl.go file.
func Index(ctx context.Context) {
	tmpl := &templates.Index{}

	// render the template
	ExecuteTemplate(ctx, tmpl)
}
