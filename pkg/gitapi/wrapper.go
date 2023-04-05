package gitapi

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/careem/githerd/pkg/file"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Wrapper interface {
	// AddRemote adds a remote to the repo.
	AddRemote(ctx context.Context, remoteName, remoteUrl string) error
	// Fetch fetches the remote.
	Fetch(ctx context.Context, remoteName string) error
	// CheckoutBranch checks out the given branch.
	CheckoutBranch(ctx context.Context, branchName string) error
	// CommitAll commits all changes.
	CommitAll(ctx context.Context, message string) error

	// Status returns the status of the repo.
	Status(ctx context.Context) error

	// CreateBranch creates a new branch.
	CreateBranch(ctx context.Context, branchName string) error
	// DeleteBranch deletes the given branch.
	DeleteBranch(ctx context.Context, branchName string) error
	// Push pushes the given branch to the remote.
	Push(ctx context.Context, remoteName, branchName string) error
	// RunCommand runs the given git command.
	RunCommand(ctx context.Context, command string, args ...string) error
}

func CloneRepository(remoteUrl, localPath, reference string) (*git.Repository, error) {
	repo, err := git.PlainClone(localPath, false, &git.CloneOptions{
		URL:           remoteUrl,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", reference)),
	})
	return repo, err
}

func MirrorRepository(remoteFetchUrl, remotePushUrl, reference string) error {
	fs := memfs.New()
	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: remoteFetchUrl,
	})
	if err != nil {
		return err
	}
	origin, err := repo.Remote("origin")
	if err != nil {
		return err
	}
	//git checkout develop
	w, err := repo.Worktree()

	ref := fmt.Sprintf("refs/heads/%s", reference)
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(ref),
		Create: true,
	})

	origin.Config().URLs = []string{remotePushUrl}

	err = origin.Push(&git.PushOptions{})
	return err
}
func LoadGitRepo(path string) (*git.Repository, error) {
	return git.PlainOpen(path)
}

func IsRepository(path string) bool {
	_, err := git.PlainOpen(path)
	//if err != nil {
	//	fmt.Println(err)
	//	return false
	//}

	if err == git.ErrRepositoryNotExists {
		return false
	}
	return err == nil
}

// FindRepositories finds all git repositories in the given folders. Goes to one level deep.
func FindRepositories(dirs []string) ([]*git.Repository, error) {
	// Iterate over the specified folders.
	repos := []*git.Repository{}
	for _, folder := range dirs {
		// Get the absolute path of the folder.
		absPath, err := filepath.Abs(folder)
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute path of folder %s: %w", folder, err)
		}
		// Check if the dir is a directory
		if !file.IsDirectory(absPath) {
			return nil, fmt.Errorf("folder %s is not a directory", absPath)
		}
		// Check if the current dir is a git repository.
		if IsRepository(absPath) {
			repo, err := LoadGitRepo(absPath)
			if err != nil {
				return nil, fmt.Errorf("failed to load git repo %s: %w", absPath, err)
			}
			repos = append(repos, repo)

		}
		// get all the subfolders.
		entries, err := os.ReadDir(absPath)
		if err != nil {
			continue
		}

		// Check if the directory is empty.
		if len(entries) == 0 {
			continue
		}
		// Iterate over the subfolders.
		for _, entry := range entries {
			// Check if the entry is a directory.
			if !entry.IsDir() {
				continue
			}
			// Check if the entry is a git repository.
			if IsRepository(filepath.Join(absPath, entry.Name())) {
				repo, err := LoadGitRepo(filepath.Join(absPath, entry.Name()))
				if err != nil {
					continue
				}
				repos = append(repos, repo)
			}
		}

	}
	return repos, nil

}
