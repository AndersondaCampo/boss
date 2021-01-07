package login

import (
	"github.com/spf13/cobra"
)

// NewCmdLogin add the command line login
func NewCmdLogin() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Register login",
		Run: func(cmd *cobra.Command, args []string) {
			login()
		},
	}
}

func login() {

}
