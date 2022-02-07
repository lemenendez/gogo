package pkg

import (
	input "github.com/lemenendez/gogo/pkg/in"
)

type Field struct {
	input.Field
	MappedType string
	IsLast     bool
}
