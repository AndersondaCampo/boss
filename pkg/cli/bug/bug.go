package bug

import (
	"github.com/hashload/boss/pkg/util"
	"github.com/spf13/cobra"
)

// NewCmdBug add the command line bug
func NewCmdBug() *cobra.Command {
	var browser bool
	cmd := &cobra.Command{
		Use:     "bugs",
		Short:   "Print package bugs",
		Long:    "Print package bugs in a web browser maybe",
		Aliases: []string{"issues"},
		Example: `  Print package bugs:
  boss bugs <pkg>

  Print package bugs in the browser:
  boss bugs <pkg> --browser`,
		Run: func(cmd *cobra.Command, args []string) {
			err := printBugs(browser)
			util.CheckErr(err)
		},
	}
	cmd.Flags().BoolVarP(&browser, "browser", "b", false, "print to the default browser")
	return cmd
}

func printBugs(browser bool) error {
	return nil
}
