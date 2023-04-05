package workspaces

// Repo represents a git repository with its absolute path and remote URL.
type Repo struct {
	Name     string `yaml:"name"`
	Path     string `yaml:"path"`
	RepoType string `yaml:"type"`
	Remote   string `yaml:"remote"`
}
