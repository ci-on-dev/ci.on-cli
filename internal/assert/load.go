package assert

import (
	"os"

	"github.com/ci-on-dev/ci.on-cli/internal/models"
	"gopkg.in/yaml.v3"
)

func LoadSuite(path string) (models.Suite, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return models.Suite{}, err
	}
	var s models.Suite
	if err := yaml.Unmarshal(data, &s); err != nil {
		return models.Suite{}, err
	}
	return s, nil
}
