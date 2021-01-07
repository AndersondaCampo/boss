package initialize

import (
	"github.com/spf13/cobra"
)

// NewCmdInitialize add the command line init
func NewCmdInitialize() *cobra.Command {
	return &cobra.Command{
		Use:     "init",
		Short:   "Initialize a new project",
		Long:    "Initialize a new project and creates a boss.json file",
		Example: "  Initialize a new project:\n  boss init\n\n  Initialize a new project without having it ask any questions:\n  boss init -q",
		Run: func(cmd *cobra.Command, args []string) {
			initializeBoss()
		},
	}
}

func initializeBoss() {

}
