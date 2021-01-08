package repo

import "github.com/spf13/cobra"

// NewCmdRepo add the command line repo
func NewCmdRepo() *cobra.Command {
	return &cobra.Command{
		Use:   "repo",
		Short: "Open package repository page in the browser",
		Example: `  Open a package:
  boss repo <pkg>`,
		Run: func(cmd *cobra.Command, args []string) {
			openRepo()
		},
	}
}

func openRepo() {

}
