package gitlab

import (
	"fmt"
	"log"
	"os"

	"github.com/ci-on-dev/ci.on-cli/internal/ir"
	"github.com/ci-on-dev/ci.on-cli/internal/models"
	"gopkg.in/yaml.v3"
)

type GitLabProvider struct{}

func NewGitLabProvider() *GitLabProvider {
	return &GitLabProvider{}
}

type GitLabWorkflow struct {
	Stages []string               `yaml:"stages"`
	Jobs   map[string]interface{} `yaml:",inline"`
}

// Parse executa gitlab-ci-local e retorna IR preenchido com jobs/steps e exit codes
func (p *GitLabProvider) Parse(suite models.Suite, test models.Test) (ir.IR, error) {
	file := suite.Pipeline.File
	// 1️⃣ Lê o arquivo
	data, err := os.ReadFile(file)
	if err != nil {
		return ir.IR{}, fmt.Errorf("cannot read file %s: %w", file, err)
	}

	// 2️⃣ Faz unmarshal no struct que detecta stages e jobs
	var wf GitLabWorkflow
	if err := yaml.Unmarshal(data, &wf); err != nil {
		return ir.IR{}, fmt.Errorf("cannot parse YAML: %w", err)
	}

	suite.Pipeline.Stages = test.Stages
	fmt.Println("Stages:", suite.Pipeline.Stages)
	wrapperFile, err := GenerateWrapperYAML(suite, test)
	if err != nil {
		log.Fatal(err)
	}

	// 3️⃣ Aqui você pode criar IR a partir do workflow
	irPipeline := ir.IR{
		SourceFile: wrapperFile,
	}

	// Adiciona jobs detectados no workflow
	for jobName := range wf.Jobs {
		j := ir.Job{
			Name: jobName,
			// Steps podem ser preenchidos depois ou apenas um step "script" por default
			Steps: []ir.Step{
				{Name: "script", Run: "placeholder", Status: "pending"},
			},
		}
		irPipeline.Jobs = append(irPipeline.Jobs, j)
	}

	// 4️⃣ Opcional: se algum job refere-se a uma stage não listada, adiciona na lista de stages
	//    (isso evita erro "stage:init not found")
	//    você pode criar um campo temporário irPipeline.Stages = wf.Stages

	return irPipeline, nil
}
