package pkg

import (
	"gopkg.in/yaml.v2"
)

func (e Entity) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(e)
}

func (e *Entity) UnmarshalYAML(b []byte) error {
	if err := yaml.Unmarshal(b, e); err != nil {
		return err
	}
	return nil
}
