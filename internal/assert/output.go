package assert

import (
	"fmt"
	"strings"

	"github.com/ci-on-dev/ci.on-cli/internal/models"
)

// Validate the output of a job contains the expected strings
func AssertOutput(a models.Assert, output string, res Result) Result {

	for _, expected := range a.Expect.OutputContains.Contains {
		if strings.Contains(output, expected) {
			// sucesso individual
			res.Assertions = append(res.Assertions, AssertionResult{
				Name:    a.Name,
				Passed:  true,
				Message: fmt.Sprintf("✅ assert passed: %s (output contains: %s)", a.Name, expected),
			})
		} else {
			// falha individual
			res.Passed = false
			res.Assertions = append(res.Assertions, NewAssetError(a.Name,
				fmt.Sprintf("(output does not contain: %q)", expected)))
		}
	}

	return res
}
