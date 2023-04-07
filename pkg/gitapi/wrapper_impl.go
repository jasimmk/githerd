package gitapi

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type gitWrapper struct {
	repo *git.Repository
}

// NewGitWrapper creates a new instance of Wrapper with the given repository.
func NewGitWrapper(path string) (Wrapper, error) {
	if !IsRepository(path) {
		return nil, fmt.Errorf("%s not a git repository", path)
	}
	repo, err := LoadGitRepo(path)
	if err != nil {
		return nil, err
	}
	return &gitWrapper{repo: repo}, err
}

func (gw *gitWrapper) AddRemote(ctx context.Context, remoteName, remoteUrl string) error {
	_, err := gw.repo.CreateRemote(&config.RemoteConfig{
		Name: remoteName,
		URLs: []string{remoteUrl},
	})
	return err
}
func (gw *gitWrapper) Fetch(ctx context.Context, remoteName string) error {
	err := gw.repo.FetchContext(ctx, &git.FetchOptions{
		RemoteName: remoteName,
	})
	return err
}
func (gw *gitWrapper) DeleteBranch(ctx context.Context, branchName string) error {
	return gw.repo.DeleteBranch(branchName)
}

func (gw *gitWrapper) CheckoutBranch(ctx context.Context, branchName string) error {
	w, err := gw.repo.Worktree()
	if err != nil {
		return err
	}
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branchName)),
	})
	return err
}

func (gw *gitWrapper) CreateBranch(ctx context.Context, branchName string) error {
	headRef, err := gw.repo.Head()
	if err != nil {
		return err
	}

	newBranchRef := plumbing.NewHashReference(plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branchName)), headRef.Hash())
	err = gw.repo.Storer.SetReference(newBranchRef)
	return err
}

func (gw *gitWrapper) Status(ctx context.Context) error {
	w, err := gw.repo.Worktree()
	if err != nil {
		return err
	}
	status, err := w.Status()
	if err != nil {
		return err
	}
	fmt.Println(status)
	return nil
}

func (gw *gitWrapper) CommitAll(ctx context.Context, message string) error {
	w, err := gw.repo.Worktree()
	if err != nil {
		return err
	}
	// Add all files
	_, err = w.Add(".")
	if err != nil {
		return err
	}
	// Commit
	_, err = w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Git Whip",
			Email: "githerd@careem.com",
			When:  time.Now(),
		},
	})
	return err
}

func (gw *gitWrapper) Push(ctx context.Context, remoteName, branchName string) error {
	err := gw.repo.PushContext(ctx, &git.PushOptions{
		RemoteName: remoteName,
		RefSpecs: []config.RefSpec{
			config.RefSpec(fmt.Sprintf("refs/heads/%s:refs/heads/%s", branchName, branchName)),
		},
	})
	return err
}

func (gw *gitWrapper) RunCommand(ctx context.Context, command string, args []string) ([]byte, error) {
	w, err := gw.repo.Worktree()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(command, args...)
	cmd.Dir = w.Filesystem.Root()
	fmt.Printf("Command:%s\n", cmd.String())
	fmt.Printf("cmd:%s, args:%s\n", command, args)
	return cmd.CombinedOutput()
}
