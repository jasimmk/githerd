package workspaces

import "github.com/careem/githerd/pkg/gitapi"

// Repo represents a git repository with its absolute path and remote URL.
type Repo struct {
	Name     string `yaml:"name"`
	Path     string `yaml:"path"`
	RepoType string `yaml:"type"`
	Remote   string `yaml:"remote"`
}

// GetGitApiWrapper returns a GitApiWrapper for the repository.
func (r *Repo) GetGitApiWrapper() (gitapi.Wrapper, error) {
	gitWrapper, err := gitapi.NewGitWrapper(r.Path)
	if err != nil {
		return nil, err
	}
	return gitWrapper, nil
}
