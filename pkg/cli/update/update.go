package update

import (
	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/hashload/boss/pkg/util"
	"github.com/spf13/cobra"
)

// NewCmdUpdate add the command line update
func NewCmdUpdate(config *configuration.Configuration) *cobra.Command {
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
			err := updateDependencies(config, noSave)
			util.CheckErr(err)
		},
	}
	cmd.Flags().BoolVar(&noSave, "no-save", false, "does not update the dependency version in the boss.json file")
	return cmd
}

func updateDependencies(config *configuration.Configuration, noSave bool) error {
	return nil
}
