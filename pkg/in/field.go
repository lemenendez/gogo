package in

type Field struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Unsigned string `yaml:"unsigned"`
	Required bool   `yaml:"required"`
	PK       bool   `yaml:"pk"`
	// Foreign Key table name
	FTable string `yaml:"ftable"`
	FKey   string `yaml:"fkey"`
	// Is unique?
	Unique bool `yaml:"unique"`
	// Size of string field
	Size    uint32 `yaml:"size"`
	Comment string `yaml:"comment"`
}
