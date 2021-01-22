package configuration

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func (c *Configuration) CurrentDir() (string, error) {
	if !c.Global {
		return os.Getwd()
	} else {
		return c.BossHome()
	}
}

func (c *Configuration) BossHome() (string, error) {
	homeDir := os.Getenv("BOSS_HOME")
	if homeDir == "" {
		systemHome, err := homedir.Dir()
		homeDir = systemHome
		if err != nil {
			return "", err
		}

		homeDir = filepath.FromSlash(homeDir)
	}
	return filepath.Join(homeDir, ".boss"), nil
}
