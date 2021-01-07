package dependencie

import (
	"github.com/spf13/cobra"
)

// NewCmdDependencie add the command line dependencies
func NewCmdDependencie() *cobra.Command {
	return &cobra.Command{
		Use:     "dependencies",
		Short:   "Print all dependencies of project",
		Aliases: []string{"dep"},
		Run: func(cmd *cobra.Command, args []string) {
			printDependencies()
		},
	}
}

func printDependencies() {

}
