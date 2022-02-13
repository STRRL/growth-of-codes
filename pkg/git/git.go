package git

type Repository interface {
	Checkout(hash string) error
	ListCommitsOfBranchOrderedByCommitTimeAsc(branch string) ([]Commit, error)
}

type RepositoryOnFileSystem interface {
	Repository
	GetRepositoryPath() string
}
