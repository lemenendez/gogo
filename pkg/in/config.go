package in

import (
	"fmt"
	"os"
)

type Config struct {
	Gens map[string]Gen `yaml:"generators,flow"`
	Name string         `yaml:"name"`
}

type Gen struct {
	Description string   `yaml:"description"`
	TplPath     string   `yaml:"template-path"`
	FieldMap    FieldMap `yaml:"map,flow"`
}

func GetDefaultConfig() *Config {
	gens := make(map[string]Gen)
	mysql := NewMySQLDDLDefault()
	gens["mysql"] = mysql
	c := &Config{
		Name: "gogo default config file",
		Gens: gens,
	}
	return c
}

func NewMySQLDDLDefault() Gen {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fm := make(FieldMap)
	fm["int"] = "int"
	fm["big_int"] = "bigint"
	fm["string"] = "varchar"
	fm["date"] = "date"
	fm["date_time"] = "datetime"
	fm["small_int"] = "smallint"
	fm["uuid"] = "binary(16)"
	fm["medium_text"] = "mediumtext"
	myddl := Gen{
		TplPath:     fmt.Sprintf("%v/gogo", home),
		FieldMap:    fm,
		Description: "MySQL DDL Generator",
	}
	return myddl
}
