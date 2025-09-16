package runner

import "github.com/ci-on-dev/ci.on-cli/internal/ir"

type DryRun struct{}

func NewDryRun() *DryRun { return &DryRun{} }

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
