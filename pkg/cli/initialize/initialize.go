package initialize

import (
	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/hashload/boss/internal/pkg/initialize"
	"github.com/hashload/boss/pkg/util"
	"github.com/spf13/cobra"
)

// NewCmdInitialize add the command line init
func NewCmdInitialize(config *configuration.Configuration) *cobra.Command {
	var quiet bool
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		Long:  "Initialize a new project and creates a boss.json file",
		Example: `  Initialize a new project:
  boss init

  Initialize a new project without having it ask any questions:
  boss init --quiet`,
		Run: func(cmd *cobra.Command, args []string) {
			initializeBoss(config, quiet)
		},
	}
	cmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "without asking questions")
	return cmd
}

func initializeBoss(config *configuration.Configuration, quiet bool) {
	err := initialize.InitalizePackage(config, quiet)
	util.CheckErr(err)
}
