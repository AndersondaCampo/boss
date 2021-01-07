package install

import (
	"github.com/spf13/cobra"
)

// NewCmdInstall add the command line install
func NewCmdInstall() *cobra.Command {
	return &cobra.Command{
		Use:     "install",
		Short:   "Install a new dependency",
		Aliases: []string{"i", "add"},
		Run: func(cmd *cobra.Command, args []string) {
			installNewDependency()
		},
	}
}

func installNewDependency() {

}
