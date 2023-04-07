package workspaces

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/jasimmk/githerd/internal/constants"
	"github.com/jasimmk/githerd/internal/gateways/reposervice"
	"github.com/jasimmk/githerd/pkg/file"
	"github.com/jasimmk/githerd/pkg/gitapi"
)

func GetWorkspacePath(name string) string {
	path, err := file.AbsPath(constants.GLOBAL_WORKSPACE_DIR)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s/%s.yaml", path, name)
}

func LoadGitRepoDataFromFolders(ctx context.Context, folders []string) ([]Repo, error) {

	gitRepos, err := gitapi.FindRepositories(folders)
	if err != nil {
		return nil, fmt.Errorf("failed to find git repositories: %w", err)
	}

	// If no repositories were found, return an error.
	if len(gitRepos) == 0 {
		return nil, errors.New("no git repositories found in the specified folders")
	}
	repos, err := getRepoData(gitRepos)
	if err != nil {
		return nil, err
	}
	return repos, nil
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
		name := file.GetFileNameFromPath(repoPath)
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
