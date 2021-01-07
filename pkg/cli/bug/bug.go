package bug

import (
	"github.com/spf13/cobra"
)

// NewCmdBug add the command line bug
func NewCmdBug() *cobra.Command {
	return &cobra.Command{
		Use:     "bugs",
		Short:   "Bugs for a package in a web browser maybe",
		Aliases: []string{"issues"},
		Run: func(cmd *cobra.Command, args []string) {
			printBugs()
		},
	}
}

func printBugs() {

}
