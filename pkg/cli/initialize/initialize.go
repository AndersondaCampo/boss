package initialize

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCmdInitialize add the command line init
func NewCmdInitialize() *cobra.Command {
	var quiet bool
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		Long:  "Initialize a new project and creates a boss.json file",
		Example: `  Initialize a new project:
  boss init

  Initialize a new project without having it ask any questions:
  boss init --quiet`,
		Run: func(cmd *cobra.Command, args []string) {
			initializeBoss(&quiet)
		},
	}
	cmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "without asking questions")
	return cmd
}

func initializeBoss(quiet *bool) {
	printHead()
}

func printHead() {
	fmt.Println(`
This utility will walk you through creating a boss.json file.
It only covers the most common items, and tries to guess sensible defaults.

Use 'boss install <pkg>' afterwards to install a package and
save it as a dependency in the boss.json file.
Press ^C at any time to quit.`)
}
