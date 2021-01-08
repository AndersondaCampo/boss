package login

import (
	"github.com/spf13/cobra"
)

// NewCmdLogin add the command line login
func NewCmdLogin() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Add a registry user account",
		Example: `  Adding a new user account:
  boss login <repo>`,
		Aliases: []string{"adduser", "add-user"},
		Run: func(cmd *cobra.Command, args []string) {
			login()
		},
	}
}

func login() {

}
