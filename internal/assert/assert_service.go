package assert

import (
	"github.com/ci-on-dev/ci.on-cli/internal/logs"
	"github.com/ci-on-dev/ci.on-cli/internal/models"
	"github.com/ci-on-dev/ci.on-cli/internal/runner"
)

type AssertService interface {
	ExecuteSuite(s models.Suite, r runner.Runner) Results
}
type assertService struct {
	logService logs.LogsService
}

func NewAssertService(logService logs.LogsService) AssertService {
	return assertService{logService: logService}
}
