package command

import "github.com/spf13/cobra"

func NewRootCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "goc-analyze",
		Short: "Analyze a Go project",
		Long:  "Analyze a Go project",
	}

	analyzeCommand, err := NewAnalyzeCommand()
	if err != nil {
		return nil, err
	}

	cmd.AddCommand(analyzeCommand)

	return cmd, nil
}
