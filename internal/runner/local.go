package runner

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ci-on-dev/ci.on-cli/internal/ir"
	logs "github.com/ci-on-dev/ci.on-cli/internal/logs"
)

type LocalRunner struct {
	logService logs.LogsService
}

func NewLocalRunner(logService logs.LogsService) *LocalRunner {
	return &LocalRunner{logService: logService}
}

type JobResult struct {
	Status string // "success" | "failed"
	Output string // stdout + stderr
}

// Run executa o pipeline localmente via gitlab-ci-local
func (r *LocalRunner) Run(p ir.IR) (map[string]JobResult, error) {
	if p.SourceFile == "" {
		return nil, fmt.Errorf("ir.SourceFile is empty")
	}

	// git add .
	gitAddCmd := exec.Command("git", "add", ".")
	gitAddCmd.Stdout = os.Stdout
	gitAddCmd.Stderr = os.Stderr
	if err := gitAddCmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run 'git add .': %w", err)
	}

	// gitlab-ci-local
	ciCmd := exec.Command("gitlab-ci-local", "--file", p.SourceFile)

	// Captura o output
	var outBuf, errBuf strings.Builder
	ciCmd.Stdout = &outBuf
	ciCmd.Stderr = &errBuf

	// adiciona variáveis
	ciCmd.Env = append(os.Environ(), fmt.Sprintf("PATH=%s", os.Getenv("PATH")))

	err := ciCmd.Run()
	status := "success"
	if err != nil || strings.Contains(outBuf.String(), "failed") || strings.Contains(errBuf.String(), "failed") {
		status = "failed"
	}

	// Preenche IR e captura output por job
	results := make(map[string]JobResult)
	for i := range p.Jobs {
		job := &p.Jobs[i]
		job.Steps[0].Status = status

		results[job.Name] = JobResult{
			Status: status,
			Output: outBuf.String() + errBuf.String(),
		}
	}

	return results, err
}
