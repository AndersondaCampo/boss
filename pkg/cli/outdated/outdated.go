package outdated

import (
	"github.com/hashload/boss/pkg/util"
	"github.com/spf13/cobra"
)

// NewCmdOutdated add the command line outdated
func NewCmdOutdated() *cobra.Command {
	return &cobra.Command{
		Use:   "outdated",
		Short: "Check for outdated packages",
		Example: `  Check all outdated packages:
  boss outdated

  Check package outdated:
  boss outdated <pkg>`,
		Run: func(cmd *cobra.Command, args []string) {
			err := listAllPackagesOutdated()
			util.CheckErr(err)
		},
	}
}

func listAllPackagesOutdated() error {
	return nil
}
