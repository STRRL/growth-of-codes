package command

import (
	"context"
	"fmt"
	"github.com/STRRL/growth-of-codes/pkg/analyze/analyzer"
	"github.com/STRRL/growth-of-codes/pkg/git"
	"github.com/STRRL/growth-of-codes/pkg/scc"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

func NewAnalyzeCommand() (*cobra.Command, error) {
	options := &AnalyzeOption{}
	cmd := &cobra.Command{
		Use:   "analyze [remote git repo url]",
		Short: "Analyze the growth of complexity of codes.",
		Long:  "Analyze the growth of complexity of codes.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("at least one argument is required")
			}
			if len(args) > 1 {
				return errors.New("only one argument is required")
			}
			return runWithOption(context.TODO(), args[0], options)
		},
	}

	cmd.Flags().StringVarP(&options.Branch, "branch", "b", "master", "branch name")
	cmd.Flags().StringVar((*string)(&options.SamplingRate), "sampling-rate", string(analyzer.PeriodWeekly), "sampling rate")

	err := cmd.RegisterFlagCompletionFunc("sampling-rate", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{
			string(analyzer.PeriodAll),
			string(analyzer.PeriodDaily),
			string(analyzer.PeriodWeekly),
			string(analyzer.PeriodMonthly),
		}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type AnalyzeOption struct {
	Branch       string
	SamplingRate analyzer.Period
	MaxPoint     uint
}

func runWithOption(ctx context.Context, repo string, options *AnalyzeOption) error {
	log.Printf("clone repository %s\n", repo)
	repository, err := git.ClonePlainGitRepository(repo, os.Stderr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("clone completed at %s\n", repository.GetRepositoryPath())
	newAnalyzer := analyzer.NewAnalyzer(
		repository,
		options.Branch,
		options.SamplingRate,
		scc.NewCommandLineSCC(),
	)
	result, err := newAnalyzer.Analyze()
	if err != nil {
		return err
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	fmt.Printf("time, complexity, commit\n")
	for _, item := range result {
		fmt.Printf("%s, %d, %s\n", item.CommitTime.In(location).Format(time.RFC3339), item.AllComplexity(), item.Commit)
	}
	return nil
}
