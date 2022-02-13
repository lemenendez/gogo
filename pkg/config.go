package pkg

import (
	"bytes"
	"text/template"
)

type Config struct {
	Gens map[string]Generator `yaml:"generators,flow"`
	Name string               `yaml:"name"`
}

type Generator struct {
	Description string   `yaml:"description"`
	FieldMap    FieldMap `yaml:"map,flow"`
	Tpl         *template.Template
}

func (g *Generator) Exec(d *Entity, tplName string) (string, error) {
	var buf bytes.Buffer
	if err := g.Tpl.ExecuteTemplate(&buf, tplName, d); err != nil {
		return "", err
	}
	return buf.String(), nil
}
