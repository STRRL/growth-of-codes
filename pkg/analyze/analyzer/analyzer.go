package analyzer

import (
	"github.com/cheggaaa/pb/v3"
	"log"
	"time"

	"github.com/STRRL/growth-of-codes/pkg/git"
	"github.com/STRRL/growth-of-codes/pkg/scc"
	"github.com/boyter/scc/v3/processor"
	"github.com/pkg/errors"
)

type Period string

const PeriodAll Period = "all"
const PeriodDaily Period = "daily"
const PeriodWeekly Period = "weekly"
const PeriodMonthly Period = "monthly"

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
	log.Println("list all commits")
	commits, err := it.repo.ListCommitsOfBranchOrderedByCommitTimeAsc(it.branch)
	if err != nil {
		return nil, err
	}

	log.Printf("sampling all %d commits with sampling rate \"%s\"\n", len(commits), it.period)
	filteredCommits := it.filteringCommitsWithPeriod(commits, it.period)

	log.Printf("analyzing %d commits\n", len(filteredCommits))

	var result []CodeComplexitySnapshot

	all := len(filteredCommits)
	bar := pb.StartNew(all)

	for _, commit := range filteredCommits {
		bar.Increment()
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
	bar.Finish()
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
		if period == PeriodDaily {
			if !it.inSameDate(commit.Time, lastDay) {
				result = append(result, commit)
				lastDay = commit.Time
			}
		}

		if period == PeriodWeekly {
			if !it.inSameWeek(commit.Time, lastDay) {
				result = append(result, commit)
				lastDay = commit.Time
			}
		}

		if period == PeriodMonthly {
			if !it.inSameMonth(commit.Time, lastDay) {
				result = append(result, commit)
				lastDay = commit.Time
			}
		}
	}
	return result
}

func (it Analyzer) inSameDate(t1 time.Time, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

func (it Analyzer) inSameWeek(t1 time.Time, t2 time.Time) bool {
	var before, after time.Time
	if t1.Before(t2) {
		before = t1
		after = t2
	} else {
		before = t2
		after = t1
	}
	return after.Sub(before).Hours() < (7 * 24)
}

func (it *Analyzer) inSameMonth(t1 time.Time, t2 time.Time) bool {
	var before, after time.Time
	if t1.Before(t2) {
		before = t1
		after = t2
	} else {
		before = t2
		after = t1
	}
	return after.Sub(before).Hours() < (30 * 24)
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
