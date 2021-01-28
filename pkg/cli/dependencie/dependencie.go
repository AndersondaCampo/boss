package dependencie

import (
	"github.com/hashload/boss/pkg/util"
	"github.com/spf13/cobra"
)

// NewCmdDependencie add the command line dependencies
func NewCmdDependencie() *cobra.Command {
	var version bool
	cmd := &cobra.Command{
		Use:     "dependencies",
		Short:   "Print all project dependencies",
		Long:    "Print all project dependencies with or without version control",
		Aliases: []string{"dep", "ls", "list", "ll", "la"},
		Example: `  Listing all dependencies:
  boss dependencies

  Listing all dependencies with version control:
  boss dependencies --version

  List package dependencies:
  boss dependencies <pkg>

  List package dependencies with version control:
  boss dependencies <pkg> --version`,
		Run: func(cmd *cobra.Command, args []string) {
			err := printDependencies(version)
			util.CheckErr(err)
		},
	}
	cmd.Flags().BoolVarP(&version, "version", "v", false, "show dependency version")
	return cmd
}

func printDependencies(version bool) error {
	return nil
}
