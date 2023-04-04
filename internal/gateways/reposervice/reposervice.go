package reposervice

import (
	"context"
)

type RepoMeta struct {
	Name      string
	Namespace string
	SshUrl    string
	WebUrl    string
}

type RepoService interface {
	// GetRepo returns a repo by its ID.
	GetRepo(ctx context.Context, id int) (*RepoMeta, error)
	// GetRepos returns all repos.
	GetRepos(ctx context.Context) ([]*RepoMeta, error)
	// GetRepoByName returns all repos with the given name.
	GetRepoByName(ctx context.Context, name string) ([]*RepoMeta, error)
}
