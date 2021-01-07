package update

import (
	"github.com/spf13/cobra"
)

// NewCmdUpdate add the command line update
func NewCmdUpdate() *cobra.Command {
	return &cobra.Command{
		Use:     "update",
		Short:   "Update dependencies",
		Aliases: []string{"up"},
		Run: func(cmd *cobra.Command, args []string) {
			updateDependencies()
		},
	}
}

func updateDependencies() {

}
