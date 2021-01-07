package boss

import (
	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/hashload/boss/pkg/cli/bug"
	"github.com/hashload/boss/pkg/cli/completion"
	"github.com/hashload/boss/pkg/cli/dependencie"
	"github.com/hashload/boss/pkg/cli/initialize"
	"github.com/hashload/boss/pkg/cli/install"
	"github.com/hashload/boss/pkg/cli/login"
	"github.com/hashload/boss/pkg/cli/owner"
	"github.com/hashload/boss/pkg/cli/uninstall"
	"github.com/hashload/boss/pkg/cli/update"
	"github.com/hashload/boss/pkg/cli/upgrade"
	"github.com/hashload/boss/pkg/cli/version"
	"github.com/spf13/cobra"
)

// InitializeBossCommandLine  initialize the boss command line
func InitializeBossCommandLine(name string) *cobra.Command {

	config := configuration.InitializeConfiguration(name)

	root := &cobra.Command{
		Use:   name,
		Short: "Dependency Manager for Delphi",
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
		},
	}

	config.BindFlags(root)

	root.AddCommand(bug.NewCmdBug())
	root.AddCommand(completion.NewCmdCompletion())
	root.AddCommand(dependencie.NewCmdDependencie())
	root.AddCommand(initialize.NewCmdInitialize())
	root.AddCommand(install.NewCmdInstall())
	root.AddCommand(login.NewCmdLogin())
	root.AddCommand(owner.NewCmdOwner())
	root.AddCommand(uninstall.NewCmdUnstall())
	root.AddCommand(update.NewCmdUpdate())
	root.AddCommand(upgrade.NewCmdUpgrade())
	root.AddCommand(version.NewCmdVersion(config))

	return root
}
