package assert

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ci-on-dev/ci.on-cli/internal/models"
)

func AssertArtifact(a models.Assert, res Result) Result {

	artifactPath := filepath.Join(
		".gitlab-ci-local",
		"artifacts",
		a.Expect.Job,
		".gitlab-ci-reports",
		"dotenv",
		a.Expect.Artifact.File,
	)

	data, err := os.ReadFile(artifactPath)
	if err != nil {
		res.Passed = false
		res.Assertions = append(res.Assertions, NewAssetError(a.Name, fmt.Sprintf("(artifact %s not found: %v)", artifactPath, err)))
	} else {
		content := string(data)
		for _, expected := range a.Expect.Artifact.Contains {
			if strings.Contains(content, expected) {
				res.Assertions = append(res.Assertions, AssertionResult{
					Name:    a.Name,
					Passed:  true,
					Message: fmt.Sprintf("✅ assert passed: %s (contains: %s)", a.Name, expected)})
			} else {
				res.Passed = false
				res.Assertions = append(res.Assertions, NewAssetError(a.Name, fmt.Sprintf("(artifact %s does not contain %q)", a.Expect.Artifact.File, expected)))

			}
		}
	}
	return res
}
