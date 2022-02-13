package pkg_test

import (
	"embed"
	"fmt"
	"github.com/lemenendez/gogo/pkg"
	"testing"
	"text/template"
)

//go:embed examples/*.*
var templateFS embed.FS

//go:embed examples/account.yml
var accYml []byte

//go:embed examples/user.yml
var usrYML []byte

//go:embed examples/role.yml
var roleYML []byte

//go:embed examples/config.yml
var configYml []byte

//go:embed examples/user-role.yml
var usrRoleYml []byte

func TestMany(t *testing.T) {
	cases := []struct {
		entitySrc []byte
		entity    *pkg.Entity
		generator string
		template  string
	}{
		{
			entitySrc: accYml,
			generator: "mysql",
			template:  "migration",
		},
		{
			entitySrc: accYml,
			generator: "mysql",
			template:  "rollback",
		},
		{
			entitySrc: accYml,
			generator: "go",
			template:  "struct",
		},
		{
			entitySrc: usrYML,
			generator: "mysql",
			template:  "migration",
		},
		{
			entitySrc: roleYML,
			generator: "mysql",
			template:  "migration",
		},
		{
			entitySrc: usrRoleYml,
			generator: "mysql",
			template:  "migration",
		},
	}

	cfg := pkg.Config{}
	if err := cfg.UnmarshalYAML(configYml); err != nil {
		t.Fatal(err)
	}
	for _, test := range cases {
		test.entity = &pkg.Entity{}
		if err := test.entity.UnmarshalYAML(test.entitySrc); err != nil {
			t.Fatal(err)
		}
		if gen, ok := cfg.Gens[test.generator]; ok {
			tpl, err := template.New("examples").Funcs(pkg.FuncMap).ParseGlob(fmt.Sprintf("%v/%v/*.tpl", "examples", test.generator))
			if err != nil {
				t.Fatal(err)
			}
			gen.Tpl = tpl
			test.entity.Map(gen.FieldMap)
			t.Log(gen.Exec(test.entity, test.template))
		} else {
			t.Fatal("invalid generator name in test case")
		}
	}
}
