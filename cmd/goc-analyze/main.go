package main

import (
	"fmt"
	"github.com/STRRL/growth-of-codes/pkg/analyze"
	"github.com/STRRL/growth-of-codes/pkg/git"
	"github.com/STRRL/growth-of-codes/pkg/scc"
	"log"
	"time"
)

func main() {
	repository, err := git.ClonePlainGitRepository("https://github.com/tikv/tikv")
	if err != nil {
		log.Fatal(err)
	}
	analyzer := analyze.NewAnalyzer(
		repository,
		"master",
		analyze.PeriodWeekly,
		scc.NewCommandLineSCC(),
	)
	result, err := analyzer.Analyze()
	if err != nil {
		log.Fatalln(err)
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("time, complexity, commit\n")
	for _, item := range result {
		fmt.Printf("%s, %d, %s\n", item.CommitTime.In(location).Format(time.RFC3339), item.AllComplexity(), item.Commit)
	}
}
