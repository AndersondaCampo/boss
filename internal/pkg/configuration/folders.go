package configuration

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// CurrentDir return a current dir
func (c *Configuration) CurrentDir() (string, error) {
	if c.Global {
		dir := os.Getenv("BOSS_HOME")
		if dir != "" {
			return dir, nil
		}
		userHome, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		return filepath.Join(userHome, ".boss"), nil
	}
	return os.Getwd()
}
