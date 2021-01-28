package version

import (
	"fmt"

	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/hashload/boss/internal/pkg/models"
	"github.com/spf13/cobra"
)

// NewCmdVersion add the command line version
func NewCmdVersion(config *configuration.Configuration) *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   fmt.Sprintf("Print the %s version", config.AppName),
		Aliases: []string{"v"},
		Example: `  Print version:
  boss version`,
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}
}

func printVersion() {
	v := models.MakeBossVersion()
	fmt.Println("Version:", v.Version)
	fmt.Println("Git commit:", v.GitCommit)
	fmt.Println("Git tree state:", v.GitTreeState)
	fmt.Println("Go version:", v.GoVersion)
}
