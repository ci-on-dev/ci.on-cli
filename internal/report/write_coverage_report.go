package report

import (
	"os"

	"gopkg.in/yaml.v3"
)

func WriteCoverageReport(report CoverageReport, path string) error {
	data, err := yaml.Marshal(&report)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
