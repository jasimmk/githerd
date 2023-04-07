package workspaces

type Config interface {
	GetRepositories() []Repo
	AddRepositories(repos []Repo)
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

// AddRepositories adds the given repositories to the workspace.
func (c *config) AddRepositories(repos []Repo) {
	for _, repo := range repos {
		if !repoContains(c.Repositories, repo) {
			c.Repositories = append(c.Repositories, repo)
		}
	}
}

func repoContains(repos []Repo, repo Repo) bool {
	for _, r := range repos {
		if r.Name == repo.Name && r.Path == repo.Path && r.Remote == repo.Remote {
			return true
		}
	}
	return false
}
