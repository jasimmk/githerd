package workspaces

import (
	"context"
	"errors"
	"fmt"

	"github.com/careem/githerd/internal/gateways/reposervice"
	"github.com/careem/githerd/pkg/filewrapper"
	"github.com/careem/githerd/pkg/gitwrapper"
	"github.com/careem/githerd/pkg/yamlwrapper"
	"github.com/go-git/go-git/v5"
)

// CreateWorkspace creates a workspace YAML file with the given name, and adds all the git repositories
// found in the specified folders to it.
func CreateWorkspace(ctx context.Context, name string, folders []string) error {
	// Create a slice to store the repositories.

	gitRepos, err := gitwrapper.FindRepositories(folders)
	if err != nil {
		return fmt.Errorf("failed to find git repositories: %w", err)
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// If no repositories were found, return an error.
	if len(gitRepos) == 0 {
		return errors.New("no git repositories found in the specified folders")
	}
	repos, err := getRepoData(gitRepos)
	if err != nil {
		return err
	}

	// Convert the slice of repositories to YAML.
	workspacePath := GetWorkspacePath(name)
	workspaceConfig := NewConfig(repos)
	err = yamlwrapper.WriteYamlFile(workspacePath, workspaceConfig)
	if err != nil {
		return err
	}

	return nil
}
func getRepoData(repos []*git.Repository) ([]Repo, error) {
	var repoData []Repo
	for _, repo := range repos {
		// Get the remote URL of the repository.
		var remoteUrl string
		remote, err := repo.Remote("origin")
		if err != nil {
			return nil, err
		}
		remoteUrl = remote.Config().URLs[0]

		// Get the absolute path of the repository.
		repoPath, err := getRepoPath(repo)
		if err != nil {
			return nil, err
		}
		// Add the repository to the slice.
		name := filewrapper.GetFileNameFromPath(repoPath)
		repoType := reposervice.DetectRemoteType(remoteUrl)
		repoData = append(repoData, Repo{
			Name:     name,
			Path:     repoPath,
			RepoType: repoType,
			Remote:   remoteUrl,
		})
	}
	return repoData, nil
}
func getRepoPath(repo *git.Repository) (string, error) {
	// Get the worktree.
	worktree, err := repo.Worktree()
	if err != nil {
		return "", err
	}

	// Get the absolute path of the repository.
	absPath := worktree.Filesystem.Root()
	return absPath, nil
}
