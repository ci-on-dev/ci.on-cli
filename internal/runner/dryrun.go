package runner

import (
	"github.com/ci-on-dev/ci.on-cli/internal/ir"
	logs "github.com/ci-on-dev/ci.on-cli/internal/logs"
)

type DryRun struct {
	logService logs.LogsService
}

func NewDryRun(logService logs.LogsService) *DryRun { return &DryRun{logService: logService} }

func (r *DryRun) Run(p ir.IR) (map[string]JobResult, error) {
	results := make(map[string]JobResult)
	for i := range p.Jobs {

		job := &p.Jobs[i]
		job.Steps[0].Status = "success"

		results[job.Name] = JobResult{
			Status: "success",
			Output: "",
		}

	}
	return results, nil
}
