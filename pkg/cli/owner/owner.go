package owner

import (
	"github.com/spf13/cobra"
)

// NewCmdOwner add the command line owner
func NewCmdOwner() *cobra.Command {
	return &cobra.Command{
		Use:     "owner",
		Short:   "Manage package owners",
		Aliases: []string{"author"},
		Run: func(cmd *cobra.Command, args []string) {
			manageOwners()
		},
	}
}

func manageOwners() {

}
