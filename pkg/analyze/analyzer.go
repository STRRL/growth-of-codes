package analyze

import (
	"github.com/STRRL/growth-of-codes/pkg/git"
	"github.com/STRRL/growth-of-codes/pkg/scc"
	"github.com/boyter/scc/v3/processor"
	"github.com/pkg/errors"
	"time"
)

type Period string

const PeriodAll Period = "All"
const PeriodDaily Period = "daily"

type Analyzer struct {
	repo   git.RepositoryOnFileSystem
	branch string
	period Period
	scc    scc.Adaptor
}

func NewAnalyzer(repo git.RepositoryOnFileSystem, branch string, period Period, scc scc.Adaptor) *Analyzer {
	return &Analyzer{repo: repo, branch: branch, period: period, scc: scc}
}

func (it *Analyzer) Analyze() ([]CodeComplexitySnapshot, error) {
	path := it.repo.GetRepositoryPath()

	commits, err := it.repo.ListCommitsOfBranchOrderedByCommitTimeAsc(it.branch)
	if err != nil {
		return nil, err
	}

	filteredCommits := it.filteringCommitsWithPeriod(commits, it.period)

	var result []CodeComplexitySnapshot

	for _, commit := range filteredCommits {
		err := it.repo.Checkout(commit.Hash)
		if err != nil {
			return nil, errors.Wrapf(err, "analyze, checkout, %s", commit.Hash)
		}
		summaries, err := it.scc.Analyze(path)
		if err != nil {
			return nil, errors.Wrapf(err, "analyze, execute scc, %s", commit.Hash)
		}
		result = append(result,
			CodeComplexitySnapshot{
				Commit:     commit.Hash,
				CommitTime: commit.Time,
				Complexity: summaries,
			})

	}

	return result, nil
}

func (it *Analyzer) filteringCommitsWithPeriod(commitsOrderByCommitTime []git.Commit, period Period) []git.Commit {
	if period == PeriodAll {
		return commitsOrderByCommitTime
	}

	if len(commitsOrderByCommitTime) == 0 {
		return nil
	}
	var result []git.Commit
	lastDay := time.Time{}

	for _, commit := range commitsOrderByCommitTime {
		if !it.isSameDate(commit.Time, lastDay) {
			result = append(result, commit)

		}
		lastDay = commit.Time
	}
	return result
}

func (it Analyzer) isSameDate(t1 time.Time, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

type CodeComplexitySnapshot struct {
	Commit     string
	CommitTime time.Time
	Complexity []processor.LanguageSummary
}

func (it *CodeComplexitySnapshot) AllComplexity() int64 {
	var result int64 = 0
	for _, item := range it.Complexity {
		result += item.Complexity
	}
	return result
}
