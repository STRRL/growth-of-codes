package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pkg/errors"
	"os"
	"reflect"
)

type PlainRepository struct {
	path string
	repo *git.Repository
}

func NewPlainRepository(path string, repo *git.Repository) *PlainRepository {
	return &PlainRepository{path: path, repo: repo}
}

func (it *PlainRepository) Checkout(hash string) error {
	worktree, err := it.repo.Worktree()
	if err != nil {
		return errors.Wrapf(err, "checkout, get worktree")
	}

	// resolve short hash
	revision, err := it.repo.ResolveRevision(plumbing.Revision(hash))
	if err != nil {
		return errors.Wrapf(err, "checkout, resolve revision for %s", hash)
	}

	err = worktree.Checkout(&git.CheckoutOptions{
		Hash:  *revision,
		Force: true,
		Keep:  false,
	})
	return nil
}

// ClonePlainGitRepository would clone the repository from the given url.
func ClonePlainGitRepository(url string) (*PlainRepository, error) {
	tempDir, err := os.MkdirTemp(os.TempDir(), "goc-git-repo-*")
	if err != nil {
		return nil, errors.Wrap(err, "create temp dir to clone the repository")
	}
	repository, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return nil, errors.Wrap(err, "clone the repository")
	}
	return NewPlainRepository(tempDir, repository), nil
}

// OpenPlainGitRepository open a git repository exist on the filesystem in the given path.
func OpenPlainGitRepository(path string) (*PlainRepository, error) {
	repository, err := git.PlainOpen(path)
	if err != nil {
		return nil, errors.Wrapf(err, " open plain repository %s", path)
	}

	return NewPlainRepository(path, repository), nil
}

func (it *PlainRepository) ListCommitsOfBranchOrderedByCommitTimeAsc(branchName string) ([]Commit, error) {
	branch, err := it.repo.Branch(branchName)
	if err != nil {
		return nil, errors.Wrapf(err, "get branch %s", branchName)
	}

	worktree, err := it.repo.Worktree()
	if err != nil {
		return nil, errors.Wrapf(err, "get worktree")
	}

	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: branch.Merge,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "checkout branch %s", branchName)
	}

	commitIter, err := it.repo.Log(&git.LogOptions{
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "get commits log of branch %s", branchName)
	}

	var result []Commit
	err = commitIter.ForEach(func(c *object.Commit) error {
		result = append(result, Commit{
			Hash: c.Hash.String(),
			Time: c.Author.When,
		},
		)
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "iterate commits log of branch %s", branchName)
	}

	it.inPlaceReverseSlice(result)
	return result, nil
}

func (it PlainRepository) inPlaceReverseSlice(commits []Commit) {
	size := reflect.ValueOf(commits).Len()
	swap := reflect.Swapper(commits)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
func (it *PlainRepository) GetRepositoryPath() string {
	return it.path
}
