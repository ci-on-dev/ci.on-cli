package gitlab

import (
	"fmt"
	"log"
	"os"

	"github.com/ci-on-dev/ci.on-cli/internal/constants"
	"github.com/ci-on-dev/ci.on-cli/internal/models"
	"gopkg.in/yaml.v3"
)

// WrapperYAML representa o YAML que será gerado dinamicamente
type WrapperYAML struct {
	Stages  []string       `yaml:"stages"`
	Vars    map[string]any `yaml:"variables,omitempty"`
	Include []string       `yaml:"include"`
}

// GenerateWrapperYAML cria um YAML temporário que inclui o workflow real e injeta variáveis
// Retorna o caminho do arquivo gerado
func GenerateWrapperYAML(suite models.Suite, test models.Test) (string, error) {
	workflowFile := suite.Pipeline.File
	if workflowFile == "" {
		return "", fmt.Errorf("workflowFile is empty")
	}

	// detecta stages automaticamente
	stages := suite.Pipeline.Stages
	if len(stages) == 0 {
		stages = []string{"init"}
	}

	// merge das variáveis do pipeline com as do teste
	vars := map[string]any{}
	for k, v := range suite.Pipeline.Vars {
		vars[k] = v
	}
	for k, v := range test.Vars {
		vars[k] = v
	}

	wrapper := WrapperYAML{
		Stages:  stages,
		Vars:    vars,
		Include: []string{workflowFile},
	}

	buf, err := yaml.Marshal(&wrapper)
	if err != nil {
		return "", fmt.Errorf("failed to marshal wrapper YAML: %w", err)
	}

	tmpDir := constants.TmpDir
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		log.Fatal(err)
	}

	// gera arquivo único por teste
	tmpFile := fmt.Sprintf("%s/ci.on.generated.%s.%s.gitlab-ci.yml",
		tmpDir,
		suite.Name,
		test.Name,
	)

	if err := os.WriteFile(tmpFile, buf, 0644); err != nil {
		return "", fmt.Errorf("failed to write wrapper YAML: %w", err)
	}

	return tmpFile, nil
}
