package configuration

import (
	"github.com/spf13/cobra"
)

type Configuration struct {
	DebugMode bool
	AppName   string
}

func InitializeConfiguration(appName string) *Configuration {
	config := &Configuration{
		DebugMode: false,
		AppName:   appName,
	}

	return config
}

func (c *Configuration) BindFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&c.DebugMode, "debug", "d", false, "Debug mode")
}
