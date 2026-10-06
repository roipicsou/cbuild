package persfile

import (
	"os"

	"gopkg.in/yaml.v3"
)

func WriteFile(nameFile string, cfg Config) error {
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(nameFile, data, 0644)
}