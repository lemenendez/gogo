package test

import (
	"github.com/lemenendez/gogo/pkg"
	"github.com/lemenendez/gogo/pkg/in"
	"testing"
)

func generate(t *testing.T, entityStr string, generatorName string, templateName string) {
	cfg, err := in.NewConfigFromFile("examples/gogo.yml")
	if err != nil {
		t.Error(err)
	}
	iEnt, err := in.NewEntityFromFile(entityStr)
	if err != nil {
		t.Error(err)
	}
	entity, err := pkg.ToEntity(*iEnt)
	if err != nil {
		t.Error(err)
	}
	if gen, ok := cfg.Gens[generatorName]; ok {
		codeGenerator, err := pkg.NewGen(gen.TplPath)
		if err != nil {
			t.Error(err)
		}
		entity.Map(gen.FieldMap)
		t.Logf("Generator %v, %v", generatorName, gen.Description)
		t.Log(codeGenerator.Generate(entity, templateName))
	}
}

func TestMyDDL1(t *testing.T) {
	generate(t, "examples/account.yml", "mysql", "migration")
	generate(t, "examples/account.yml", "go", "entity")
}
