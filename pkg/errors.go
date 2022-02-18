package pkg

import "errors"

var (
	ErrFieldTypeNotExists error = errors.New("entity type does not exist in the generator field map")
)
