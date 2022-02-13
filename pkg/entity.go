package pkg

type Entity struct {
	Name    string   `yaml:"name"`
	Comment string   `yaml:"comment"`
	Fields  []Field  `yaml:"fields,flow"`
	Props   FieldMap `yaml:"props"`
	PK      *Field
}

type Field struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Unsigned bool   `yaml:"unsigned"`
	Required bool   `yaml:"required"`
	PK       bool   `yaml:"pk"`
	// Foreign Key table name
	FTable string `yaml:"ftable"`
	FKey   string `yaml:"fkey"`
	// Is unique?
	Unique bool `yaml:"unique"`
	// Size of string field
	Size       uint32 `yaml:"size"`
	Comment    string `yaml:"comment"`
	MappedType string
	IsLast     bool
}

// Map maps fields-types from a FieldMap
func (e *Entity) Map(fm FieldMap) {
	for idx, f := range e.Fields {
		e.Fields[idx].MappedType = fm[f.Type]
		if f.PK {
			e.PK = &e.Fields[idx]
		}
	}
}
