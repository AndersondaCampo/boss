package boss

import (
	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/hashload/boss/pkg/cli/bug"
	"github.com/hashload/boss/pkg/cli/completion"
	"github.com/hashload/boss/pkg/cli/dependencie"
	"github.com/hashload/boss/pkg/cli/initialize"
	"github.com/hashload/boss/pkg/cli/install"
	"github.com/hashload/boss/pkg/cli/login"
	"github.com/hashload/boss/pkg/cli/logout"
	"github.com/hashload/boss/pkg/cli/outdated"
	"github.com/hashload/boss/pkg/cli/owner"
	"github.com/hashload/boss/pkg/cli/repo"
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
	root.AddCommand(initialize.NewCmdInitialize(config))
	root.AddCommand(install.NewCmdInstall(config))
	root.AddCommand(login.NewCmdLogin())
	root.AddCommand(logout.NewCmdLogout())
	root.AddCommand(outdated.NewCmdOutdated())
	root.AddCommand(owner.NewCmdOwner())
	root.AddCommand(repo.NewCmdRepo())
	root.AddCommand(uninstall.NewCmdUnstall())
	root.AddCommand(update.NewCmdUpdate())
	root.AddCommand(upgrade.NewCmdUpgrade())
	root.AddCommand(version.NewCmdVersion(config))

	return root
}
