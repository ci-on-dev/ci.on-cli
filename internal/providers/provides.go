package providers

import (
	"github.com/ci-on-dev/ci.on-cli/internal/ir"
	"github.com/ci-on-dev/ci.on-cli/internal/models"
	"github.com/ci-on-dev/ci.on-cli/internal/providers/gitlab"
)

type Provider interface {
	Parse(suite models.Suite, test models.Test) (ir.IR, error)
}

func Get(name string) Provider {
	switch name {
	case "github":
		return &GitHubProvider{}
	case "gitlab":
		return &gitlab.GitLabProvider{}
	case "azdo":
		return &AzDOProvider{}
	default:
		return &GitHubProvider{} // fallback
	}
}

// GitHubProvider is a stub implementation of the Provider interface.

type AzDOProvider struct{}

func (p *AzDOProvider) Parse(suite models.Suite, test models.Test) (ir.IR, error) {
	return ir.IR{}, nil
}
