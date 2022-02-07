package pkg

import (
	input "github.com/lemenendez/gogo/pkg/in"
)

type Entity struct {
	input.Entity
	Fields []Field
	PK     *Field
}

// Map maps fields from a FieldMap
func (e *Entity) Map(fm input.FieldMap) {
	for idx, f := range e.Fields {
		e.Fields[idx].MappedType = fm[f.Type]
	}
	// TODO: why do I need to do this?
	if e.PK != nil {
		e.PK.MappedType = fm[e.PK.Type]
	}
}
