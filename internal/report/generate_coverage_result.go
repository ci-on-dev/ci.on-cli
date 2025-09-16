package report

import "github.com/ci-on-dev/ci.on-cli/internal/assert"

func GenerateCoverageFromResult(suiteName string, pipelineFile string, res assert.Result) CoverageReport {
	jobCov := JobCoverage{
		Name:   res.Suite,
		Status: "success", // ou "failed"
	}

	// for _, a := range res.Assertions {
	// 	assertCov := AssertCoverage{
	// 		Name:   a.Name,
	// 		Result: map[bool]string{true: "passed", false: "failed"}[a.Passed],
	// 	}

	// 	if a.Artifact != "" {
	// 		assertCov.Artifact = a.Artifact
	// 		assertCov.Checks = make(map[string]string)
	// 		for val, ok := range a.Contains {
	// 			if ok {
	// 				assertCov.Checks[val] = "passed"
	// 			} else {
	// 				assertCov.Checks[val] = "failed"
	// 			}
	// 		}
	// 	}

	// 	jobCov.Asserts = append(jobCov.Asserts, assertCov)
	// }

	return CoverageReport{
		Pipelines: []PipelineCoverage{
			{
				File: pipelineFile,
				Jobs: []JobCoverage{jobCov},
			},
		},
	}
}
