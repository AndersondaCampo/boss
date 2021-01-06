package util

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func TryGetEnv(name string) string {
	if value, found := os.LookupEnv(name); found {
		return value
	} else {
		logrus.Warnf("Variable %s not found on env", name)
		return ""
	}
}

func GetEnvDef(name string, def string) string {
	if value, found := os.LookupEnv(name); found {
		return value
	} else {
		logrus.Warnf("Variable %s not found on env", name)
		return def
	}
}

func GetEnvOrFatal(name string) string {
	if value, found := os.LookupEnv(name); found {
		return value
	} else {
		CheckErr(fmt.Errorf("Variable %s not found on env", name))
		return ""
	}
}
