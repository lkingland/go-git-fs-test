package gitfs

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Clone the repository into memory and return HEAD as a billy Filesystem.
func Clone(url string) (billy.Filesystem, error) {
	r, err := git.Clone(
		memory.NewStorage(),
		memfs.New(),
		&git.CloneOptions{
			URL:               url,
			Depth:             1,
			Tags:              git.NoTags,
			RecurseSubmodules: git.NoRecurseSubmodules,
		})
	if err != nil {
		return nil, err
	}
	wt, err := r.Worktree()
	if err != nil {
		return nil, err
	}
	return wt.Filesystem, nil
}
