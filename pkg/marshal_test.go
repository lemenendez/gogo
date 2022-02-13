package pkg_test

import (
	"github.com/lemenendez/gogo/pkg"
	"testing"
)

func TestConfigMarshal(t *testing.T) {
	cfg := &pkg.Config{
		Name: "my configuration",
	}
	cfg.Gens = make(map[string]pkg.Generator)
	data, err := cfg.MarshalYAML()
	if err != nil {
		t.Failed()
	}
	err = cfg.UnmarshalYAML(data)
	if err != nil {
		t.Failed()
	}
}

func TestConfigGenMarshal(t *testing.T) {
	cfg := &pkg.Config{
		Name: "my configuration",
	}
	cfg.Gens = make(map[string]pkg.Generator)
	gen := pkg.Generator{Description: "some conde generator"}
	cfg.Gens["gen1"] = gen
	data, err := cfg.MarshalYAML()
	if err != nil {
		t.Failed()
	}
	err = cfg.UnmarshalYAML(data)
	if err != nil {
		t.Failed()
	}
}

func TestConfigMarshalError(t *testing.T) {
	data := make([]byte, 10, 10)
	cfg := &pkg.Config{}
	err := cfg.UnmarshalYAML(data)
	if err == nil {
		t.Failed()
	}
}

func TestEntityMarshal(t *testing.T) {
	ent := pkg.Entity{
		Name: "user",
	}

	data, err := ent.MarshalYAML()
	if err != nil {
		t.Failed()
	}
	err = ent.UnmarshalYAML(data)
	if err != nil {
		t.Failed()
	}
}

func TestEntityMarshalError(t *testing.T) {
	data := make([]byte, 10, 10)
	e := &pkg.Entity{}
	err := e.UnmarshalYAML(data)
	if err == nil {
		t.Failed()
	}
}
