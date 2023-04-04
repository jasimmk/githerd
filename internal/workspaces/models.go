package workspaces

type Workspace interface {
	// GetRepoFolders returns the absolute paths of all repositories in the workspace.
	GetRepoFolders(repoType string) []string
	GetName() string
	GetFilePath() string
	GetConfig() Config
}

type workspace struct {
	Name     string
	FilePath string
	Config   Config
}

func NewWorkspace(name, path string, config Config) Workspace {
	return &workspace{
		Name:     name,
		FilePath: path,
		Config:   config,
	}
}

// GetRepoFolders returns the absolute paths of all repositories in the workspace.
// Provide the type of repository to filter the results, or an empty string to return all repositories.
func (w *workspace) GetRepoFolders(repoType string) []string {
	var checkRepoType bool
	if repoType != "" {
		checkRepoType = true
	}

	var repoFolders []string
	for _, repo := range w.Config.GetRepositories() {
		if checkRepoType && repo.RepoType == repoType {
			repoFolders = append(repoFolders, repo.Path)
		} else {
			repoFolders = append(repoFolders, repo.Path)
		}
	}
	return repoFolders
}

// GetName() returns the name of the workspace.
func (w *workspace) GetName() string {
	return w.Name
}

// GetFilePath() returns the path of the workspace file.
func (w *workspace) GetFilePath() string {
	return w.FilePath
}

// GetConfig() returns the workspace configuration.
func (w *workspace) GetConfig() Config {
	return w.Config
}
