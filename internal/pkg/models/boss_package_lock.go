package models

import "time"

type bossDependencyArtifacts struct {
	Bin []string `json:"bin,omitempty"`
	Dcp []string `json:"dcp,omitempty"`
	Dcu []string `json:"dcu,omitempty"`
	Bpl []string `json:"bpl,omitempty"`
}

type bossLockedDependency struct {
	Name      string                  `json:"name"`
	Version   string                  `json:"version"`
	Hash      string                  `json:"hash"`
	Artifacts bossDependencyArtifacts `json:"artifacts"`
	Failed    bool                    `json:"failed"`
	Changed   bool                    `json:"changed"`
}

// BossPackageLock is a model of boss-lock.json
type BossPackageLock struct {
	fileName  string
	Hash      string                          `json:"hash"`
	Updated   time.Time                       `json:"updated"`
	Installed map[string]bossLockedDependency `json:"installedModules"`
}
