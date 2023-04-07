package workspaces

import (
	"context"
	"fmt"
)

// UpsertWorkspace adds the given repositories to the workspace, or creates a new workspace if it doesn't exist.
func UpsertWorkspace(ctx context.Context, name string, folders []string) error {

	workspace, err := LoadWorkspace(name)
	// Workspace exists
	if err == nil {
		repos, err := LoadGitRepoDataFromFolders(ctx, folders)
		if err != nil {
			return err
		}
		workspace.GetConfig().AddRepositories(repos)
		err = Save(name, workspace.GetConfig())
		if err == nil {
			fmt.Printf("Workspace:%s updated successfully\n", name)
		}

		return err
	} else {
		// If the workspace doesn't exist, create it.
		err = CreateWorkspace(ctx, name, folders)
		if err == nil {
			fmt.Printf("Workspace:%s created successfully\n", name)
		}
		return err

	}
}
