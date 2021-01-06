package boss

import (
	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/hashload/boss/pkg/cli/completion"
	"github.com/hashload/boss/pkg/cli/version"
	"github.com/spf13/cobra"
)

func NewBossCommand(name string) *cobra.Command {

	config := configuration.InitializeConfiguration(name)

	root := &cobra.Command{
		Use:   name,
		Short: "Dependency Manager for Delphi",
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
		},
	}

	config.BindFlags(root)

	root.AddCommand(version.NewCmdVersion(config))
	root.AddCommand(completion.NewCmdCompletion())

	return root
}
