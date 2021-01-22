package initialize

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/hashload/boss/internal/pkg/models"
	"github.com/hashload/boss/internal/pkg/ui"
)

func InitalizePackage(config *configuration.Configuration, quiet bool) error {
	folderName := ""
	wd, err := config.CurrentDir()
	if err != nil {
		return err
	}
	var finalFile = filepath.Join(wd, "boss.json")

	if stat, err := os.Stat(finalFile); err == nil && !stat.IsDir() {
		if ok, _ := ui.GetConfirmation("boss.json already exists, do you want to overwrite it", false); !ok {
			return nil
		}
	}

	rxp, err := regexp.Compile(`^.+\` + string(filepath.Separator) + `([^\\]+)$`)
	if err == nil {
		allString := rxp.FindAllStringSubmatch(wd, -1)
		folderName = allString[0][1]
	}

	var pkg = models.MakeBossPackage()
	if quiet {
		pkg.Name = folderName
		pkg.Version = "1.0.0"
		pkg.MainSrc = "./"
	} else {
		pkg.Name = ui.GetTextOrDef("Package name", folderName)
		pkg.Homepage = ui.GetTextOrDef("homepage", "")
		pkg.Version = ui.GetTextOrDef("version: (1.0.0)", "1.0.0")
		pkg.Description = ui.GetTextOrDef("description", "")
		pkg.MainSrc = ui.GetTextOrDef("source folder: (./)", "./")
	}

	return pkg.SaveToFile(finalFile)
}
