package assert

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ci-on-dev/ci.on-cli/internal/models"
	"github.com/ci-on-dev/ci.on-cli/internal/providers"
	"github.com/ci-on-dev/ci.on-cli/internal/runner"
)

type AssertionResult struct {
	Name    string // nome do assert
	Passed  bool
	Message string // mensagem de sucesso ou erro
}

func (service assertService) ExecuteSuite(s models.Suite, r runner.Runner) Results {
	fmt.Println("Running suite:", s.Name)
	var results Results
	for _, test := range s.Tests {
		fmt.Printf("\n=== Running test: %s ===\n", test.Name)
		service.logService.Debugf("Provider: %s", s.Pipeline.Provider)
		prov := providers.Get(s.Pipeline.Provider)
		if prov == nil {
			log.Fatal("provider not found")
		}
		service.logService.Debugf("Pipeline file: %s", s.Pipeline.File)
		irPipeline, err := prov.Parse(s, test)
		if err != nil {
			log.Fatal(err)
		}
		service.logService.Debugf("IR Pipeline: %s", irPipeline)
		runResult, err := r.Run(irPipeline)
		service.logService.Debugf("Run Result: %s", runResult)
		if err != nil {
			return Results{
				Results: []Result{
					{Suite: s.Name,
						Passed: false,
						Assertions: []AssertionResult{
							{
								Name:    "runner error",
								Passed:  false,
								Message: fmt.Sprintf("runner error: %v", err),
							},
						},
					},
				},
			}
		}

		fmt.Println("Asserting results...")
		res := Result{
			Suite:  s.Name,
			Test:   test.Name,
			Passed: err == nil,
		}

		for _, a := range test.Asserts {
			fmt.Printf("\n=== Assert: %s ===\n", a.Name)
			// job := ir.FindJob(a.Expect.Job)
			// if job == nil {
			// 	res.Passed = false
			// 	res.Assertions = append(res.Assertions, NewAssetError(a.Name, fmt.Sprintf("(job %s not found)", a.Expect.Job)))
			// 	continue
			// }

			// if a.Expect.Status != "" {
			// 	jobRes, ok := results[a.Expect.Job]
			// 	if !ok || jobRes.Status != a.Expect.Status {
			// 		res.Passed = false
			// 		got := "not found"
			// 		if ok {
			// 			got = jobRes.Output
			// 		}
			// 		res.Assertions = append(res.Assertions, NewAssetError(a.Name, fmt.Sprintf("(expected %s got %s)", a.Expect.Status, got)))
			// 		continue
			// 	}
			// }

			// if a.Expect.Step != "" {
			// 	stepFound := false
			// 	for _, s := range job.Steps {
			// 		if s.Name == a.Expect.Step {
			// 			stepFound = true
			// 			if a.Expect.StepStatus != "" && s.Status != a.Expect.StepStatus {
			// 				res.Passed = false
			// 				res.Assertions = append(res.Assertions, NewAssetError(a.Name, fmt.Sprintf("(step %s status expected %s got %s)", s.Name, a.Expect.StepStatus, s.Status)))
			// 			}
			// 			if s.Status != a.Expect.StepStatus {
			// 				res.Passed = false
			// 				res.Assertions = append(res.Assertions, NewAssetError(a.Name, fmt.Sprintf("(step %s status expected %s got %s)", s.Name, a.Expect.StepStatus, s.Status)))
			// 			}
			// 			break
			// 		}
			// 	}
			// 	if !stepFound {
			// 		res.Passed = false
			// 		res.Assertions = append(res.Assertions, NewAssetError(a.Name, fmt.Sprintf("(step %s not found)", a.Expect.Step)))
			// 	}
			// }

			if a.Expect.Artifact != nil {
				res = AssertArtifact(a, res)
			}

			if a.Expect.OutputContains != nil {
				if len(a.Expect.OutputContains.Contains) > 0 {
					output := runResult[a.Expect.Job].Output
					res = AssertOutput(a, output, res)
				}
			}
		}

		for _, e := range res.Assertions {
			fmt.Println(e.Message)
		}

		results.Results = append(results.Results, res)
		if !res.Passed {
			results.Passed = false
		}

		if len(test.After) > 0 {
			for _, after := range test.After {
				cmd := exec.Command("sh", "-c", after)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				service.logService.Debugf("running after: %s", after)
				if err := cmd.Run(); err != nil {
					service.logService.Debugf("after error: %v", err.Error())
				}
			}
		}
	}

	return results
}

func NewAssetError(name, message string) AssertionResult {

	return AssertionResult{
		Name:    name,
		Passed:  false,
		Message: fmt.Sprintf("❌ assert failed: %s %s", name, message)}

}
