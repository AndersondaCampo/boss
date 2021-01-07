package initialize

import (
	"github.com/spf13/cobra"
)

// NewCmdInitialize add the command line init
func NewCmdInitialize() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		Run: func(cmd *cobra.Command, args []string) {
			initializeBoss()
		},
	}
}

func initializeBoss() {

}
