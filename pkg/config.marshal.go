package pkg

import (
	"gopkg.in/yaml.v2"
)

func (cfg Config) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(cfg)
}

func (cfg *Config) UnmarshalYAML(b []byte) error {
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return err
	}
	return nil
}
