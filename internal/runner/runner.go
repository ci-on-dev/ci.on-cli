package runner

import "github.com/ci-on-dev/ci.on-cli/internal/ir"

type Runner interface {
	Run(p ir.IR) (map[string]JobResult, error)
}

type Config struct {
	Network string
}
