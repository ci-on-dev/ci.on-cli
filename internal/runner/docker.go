package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ci-on-dev/ci.on-cli/internal/ir"
	logs "github.com/ci-on-dev/ci.on-cli/internal/logs"
)

type Docker struct {
	cfg        Config
	logService logs.LogsService
}

func NewDocker(cfg Config, logService logs.LogsService) *Docker {
	return &Docker{cfg: cfg, logService: logService}
}

// Run executa cada step do job dentro do host usando docker run
func (d *Docker) Run(p ir.IR) (map[string]JobResult, error) {
	results := make(map[string]JobResult)
	for i := range p.Jobs {
		job := &p.Jobs[i]
		jobStatus := "success"

		fmt.Println("=== Running job:", job.Name, "===")

		for j := range job.Steps {
			step := &job.Steps[j]
			fmt.Println("-> Step:", step.Name)

			// monta o comando Docker para executar o step
			cmd := exec.Command("docker", "run",
				"--rm",
				"-v", fmt.Sprintf("%s:/workspace", os.Getenv("PWD")), // monta workspace
				"-w", "/workspace",
				"golang:1.25", // imagem golang oficial
				"sh", "-c", step.Run,
			)

			out, err := cmd.CombinedOutput()
			step.Logs = append(step.Logs, string(out))

			if err != nil {
				step.Status = "failed"
				jobStatus = "failed"
				fmt.Println("❌ Step failed:", step.Name)
			} else {
				step.Status = "success"
				fmt.Println("✅ Step succeeded:", step.Name)
			}
			results[job.Name] = JobResult{
				Status: jobStatus,
				Output: string(out),
			}
		}
	}

	return results, nil
}
