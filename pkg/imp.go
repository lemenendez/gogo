package pkg

import (
	"bytes"
	"fmt"
	"text/template"
)

func (g *Generator) renderTpl(d interface{}, tplName string) string {
	var buf bytes.Buffer
	g.tpl.ExecuteTemplate(&buf, tplName, d)
	return buf.String()
}

func (g *Generator) Generate(d *Entity, template string) string {
	return g.renderTpl(d, template)
}

func NewGen(templateFolder string) (*Generator, error) {
	tpl, err := template.New("principal-main-template-long-name").Funcs(FuncMap).ParseGlob(fmt.Sprintf("%v/*.tpl", templateFolder))
	if err != nil {
		return nil, err
	}
	tpl.Funcs(FuncMap)
	g := &Generator{tpl: tpl}
	return g, err
}
