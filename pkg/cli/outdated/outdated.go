package outdated

import "github.com/spf13/cobra"

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
			listAllPackagesOutdated()
		},
	}
}

func listAllPackagesOutdated() {

}
