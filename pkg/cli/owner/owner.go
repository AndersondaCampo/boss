package owner

import (
	"github.com/hashload/boss/pkg/util"
	"github.com/spf13/cobra"
)

// NewCmdOwner add the command line owner
func NewCmdOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owner",
		Short: "Manage package owners",
		Long: `Manage ownership of published packages:
ls: List all the users who have access to modify a package and push new versions. Handy when you need to know who to bug for help.
add: Add a new user as a maintainer of a package. This user is enabled to modify metadata, publish new versions, and add other owners.
rm: Remove a user from the package owner list. This immediately revokes their privileges.`,
		Aliases: []string{"author"},
		Example: `  boss owner add <user>
  boss owner rm <user>
  boss owner ls`,
		Run: func(cmd *cobra.Command, args []string) {
			err := manageOwners()
			util.CheckErr(err)
		},
	}
	return cmd
}

func manageOwners() error {
	return nil
}
