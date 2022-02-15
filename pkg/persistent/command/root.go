package command

import "github.com/spf13/cobra"

func NewRootCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "goc-persistent",
		Short: "Persistent is a tool for uploading goc-analyzed data into database. No common purpose used.",
	}

	return cmd, nil
}
