package providers

import (
	"os"

	"github.com/ci-on-dev/ci.on-cli/internal/ir"
	"github.com/ci-on-dev/ci.on-cli/internal/models"
	"gopkg.in/yaml.v3"
)

type GitHubProvider struct{}

func (g *GitHubProvider) Parse(suite models.Suite, test models.Test) (ir.IR, error) {
	data, err := os.ReadFile(suite.Pipeline.File)
	if err != nil {
		return ir.IR{}, err
	}

	// estrutura simplificada do workflow
	var wf struct {
		Jobs map[string]struct {
			Steps []struct {
				Name string `yaml:"name"`
				Run  string `yaml:"run,omitempty"`
				Uses string `yaml:"uses,omitempty"`
			} `yaml:"steps"`
		} `yaml:"jobs"`
	}

	if err := yaml.Unmarshal(data, &wf); err != nil {
		return ir.IR{}, err
	}

	irPipeline := ir.IR{}
	for jobName, job := range wf.Jobs {
		j := ir.Job{Name: jobName}
		for _, step := range job.Steps {
			s := ir.Step{
				Name: step.Name,
				Run:  step.Run,
			}
			if step.Uses != "" {
				s.Run = "uses: " + step.Uses
			}
			j.Steps = append(j.Steps, s)
		}
		irPipeline.Jobs = append(irPipeline.Jobs, j)
	}
	return irPipeline, nil
}
