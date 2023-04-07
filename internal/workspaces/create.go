package workspaces

import (
	"context"
)

// CreateWorkspace creates a workspace YAML file with the given name, and adds all the git repositories
// found in the specified folders to it.
func CreateWorkspace(ctx context.Context, name string, folders []string) error {
	repos, err := LoadGitRepoDataFromFolders(ctx, folders)
	if err != nil {
		return err
	}
	// Convert the slice of repositories to YAML.
	workspaceConfig := NewConfig(repos)

	return Save(name, workspaceConfig)
}
