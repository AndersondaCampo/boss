package update

import (
	"github.com/spf13/cobra"
)

// NewCmdUpdate add the command line update
func NewCmdUpdate() *cobra.Command {
	var noSave bool
	cmd := &cobra.Command{
		Use:     "update",
		Short:   "Update dependencies",
		Aliases: []string{"up"},
		Example: `  Update a package:
  boss update <pkg>

  Update all packages:
  boss update

  Update a package without update it from the boss.json file:
  boss update <pkg> --no-save`,
		Run: func(cmd *cobra.Command, args []string) {
			updateDependencies()
		},
	}
	cmd.Flags().BoolVar(&noSave, "no-save", false, "does not update the dependency version in the boss.json file")
	return cmd
}

func updateDependencies() {

}
