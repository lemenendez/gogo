package in

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
)

func UnmarshalConfig(data []byte) (*Config, error) {
	var c Config
	if err := yaml.UnmarshalStrict(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func UnmarshalEntity(data []byte) (*Entity, error) {
	var d Entity
	if err := yaml.UnmarshalStrict(data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func NewConfigFromFile(filename string) (*Config, error) {
	bData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cfg, err := UnmarshalConfig(bData)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}

func NewEntityFromFile(filename string) (*Entity, error) {
	bData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ent, err := UnmarshalEntity(bData)
	if err != nil {
		panic(err)
	}
	return ent, nil
}
func SaveDefaultConfig(filename string) error {
	exists, err := FileExists(filename)
	if err != nil {
		return err
	}
	if !exists {
		cfg := GetDefaultConfig()
		bdata, err := yaml.Marshal(cfg)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, bdata, 0644)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
