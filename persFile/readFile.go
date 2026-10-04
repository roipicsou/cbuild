package persfile

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFile(nameFile string) (*Config, error) {
	data, err := os.ReadFile(nameFile)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}