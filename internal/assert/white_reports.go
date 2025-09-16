package assert

import (
	"fmt"

	"github.com/ci-on-dev/ci.on-cli/internal/models"
)

func (r Results) WriteReports(reports []models.Report) error {
	// TODO: implementar outros formatos (junit, json, etc)
	// por enquanto só imprime no console
	for _, rep := range reports {
		fmt.Printf("report: would write %s to %s\n", rep.Type, rep.Path)
	}
	return nil
}
