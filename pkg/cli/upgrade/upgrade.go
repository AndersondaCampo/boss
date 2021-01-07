package upgrade

import (
	"github.com/spf13/cobra"
)

// NewCmdUpgrade add the command line upgrage
func NewCmdUpgrade() *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade the client version",
		Run: func(cmd *cobra.Command, args []string) {
			upgradeBoss()
		},
	}
}

func upgradeBoss() {

}
