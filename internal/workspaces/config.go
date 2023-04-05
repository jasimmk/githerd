package workspaces

type Config interface {
	GetRepositories() []Repo
}
type config struct {
	Repositories []Repo `yaml:"repositories"`
}

// NewConfig creates a new workspace configuration.
func NewConfig(repos []Repo) Config {
	return &config{
		Repositories: repos,
	}
}

// GetRepositories returns the repositories in the workspace.
func (c *config) GetRepositories() []Repo {
	return c.Repositories
}
