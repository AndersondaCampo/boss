package uninstall

import (
	"github.com/spf13/cobra"
)

// NewCmdUnstall add the command line uninstall
func NewCmdUnstall() *cobra.Command {
	var noSave bool
	cmd := &cobra.Command{
		Use:     "uninstall",
		Short:   "Uninstall a dependency",
		Long:    "This uninstalls a package, completely removing everything boss installed on its behalf",
		Aliases: []string{"remove", "rm", "r", "un", "unlink"},
		Example: `  Uninstall a package:
  boss uninstall <pkg>

  Uninstall a package without removing it from the boss.json file:
  boss uninstall <pkg> --no-save`,
		Run: func(cmd *cobra.Command, args []string) {
			uninstallDependency(&noSave)
		},
	}
	cmd.Flags().BoolVar(&noSave, "no-save", false, "package will not be removed from your boss.json file")
	return cmd
}

func uninstallDependency(noSave *bool) {

}
