package render

import (
	"fmt"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/unrolled/render"
	"html/template"
	"io"
)

type Client interface {
	Render(w io.Writer, template string, data interface{})
}

type Config struct {
	TemplatePath string
}

type clientImpl struct {
	engine   *render.Render
	template *template.Template
}

func New(cfg Config) Client {
	//templates := template.Must(template.ParseGlob(fmt.Sprint(cfg.TemplatePath, "template/section/*.html")))
	templates := template.Must(template.ParseGlob(fmt.Sprint(cfg.TemplatePath, "template/*.html")))
	fmt.Printf("list template, %v", templates)
	return &clientImpl{
		engine: render.New(render.Options{
			Directory: "files/template",
		}),
		template: templates,
	}
}

func (c *clientImpl) Render(w io.Writer, template string, data interface{}) {
	//c.engine.HTML(w, http.StatusOK, template, data)
	tmpl := c.template.Lookup(template)
	fmt.Printf("retrieve template: %v", tmpl)
	if tmpl == nil {
		log.WarnDetail("RENDER", "warning template not found")
		return
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.ErrorDetail("RENDER", "error while render, Err: %v", err)
	}
}
