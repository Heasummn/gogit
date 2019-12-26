package util

import (
	"os"

	"gopkg.in/src-d/go-git.v4"
)

type GitInfo struct {
	repo *git.Repository
}

// InitGitInfo returns true if repo exists
func InitGitInfo(g *GitInfo) bool {
	path, _ := os.Getwd()
	repo, err := git.PlainOpen(path)
	g.repo = repo

	if err == git.ErrRepositoryNotExists {
		return false
	}
	return true
}

func (g *GitInfo) GetStaging() git.Status {
	worktree, err := g.repo.Worktree()
	if err != nil {
		panic(err)
	}
	status, _ := worktree.Status()
	return status
}
