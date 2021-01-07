package uninstall

import (
	"github.com/spf13/cobra"
)

// NewCmdUnstall add the command line uninstall
func NewCmdUnstall() *cobra.Command {
	return &cobra.Command{
		Use:     "uninstall",
		Short:   "Uninstall a dependency",
		Aliases: []string{"remove", "rm", "r", "un", "unlink"},
		Run: func(cmd *cobra.Command, args []string) {
			uninstallDependency()
		},
	}
}

func uninstallDependency() {

}
