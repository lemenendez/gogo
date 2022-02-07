package pkg

import (
	input "github.com/lemenendez/gogo/pkg/in"
)

// ToEntity converts a struct of pkg.in.Entity type to pkg.Entity type
func ToEntity(e input.Entity) (*Entity, error) {
	fields := make([]Field, 0)
	var pk *Field
	for i := 0; i < len(e.Fields); i++ {
		field := Field{
			Field: e.Fields[i],
		}
		if e.Fields[i].PK {
			pk = &field
		}
		fields = append(fields, field)
	}
	if len(fields) > 0 {
		fields[len(fields)-1].IsLast = true
	}
	ent := &Entity{
		Entity: e,
		Fields: fields,
		PK:     pk,
	}
	return ent, nil
}
