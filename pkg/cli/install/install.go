package install

import (
	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/hashload/boss/pkg/util"
	"github.com/spf13/cobra"
)

// NewCmdInstall add the command line install
func NewCmdInstall(config *configuration.Configuration) *cobra.Command {
	var noSave bool
	cmd := &cobra.Command{
		Use:     "install",
		Short:   "Install a new dependency",
		Aliases: []string{"i", "add"},
		Example: `  Add a new dependency:
  boss install <pkg>

  Add a new version-specific dependency:
  boss install <pkg>@<version>

  Install a dependency without add it from the boss.json file:
  boss install <pkg> --no-save`,
		Run: func(cmd *cobra.Command, args []string) {
			err := installDependency(config, noSave)
			util.CheckErr(err)
		},
	}
	cmd.Flags().BoolVar(&noSave, "no-save", false, "prevents saving to `dependencies`")
	return cmd
}

func installDependency(config *configuration.Configuration, noSave bool) error {
	return nil
}
