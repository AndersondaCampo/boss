package logout

import (
	"github.com/hashload/boss/pkg/util"
	"github.com/spf13/cobra"
)

// NewCmdLogout add the command line logout
func NewCmdLogout() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Log out of the registry",
		Example: `  Remove a registry user account:
  boss logout <repo>`,
		Run: func(cmd *cobra.Command, args []string) {
			err := logout()
			util.CheckErr(err)
		},
	}
}

func logout() error {
	return nil
}
