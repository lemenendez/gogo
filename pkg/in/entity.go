package in

type Entity struct {
	Name    string   `yaml:"name"`
	Comment string   `yaml:"comment"`
	Fields  []Field  `yaml:"fields,flow"`
	Props   FieldMap `yaml:"props"`
}
