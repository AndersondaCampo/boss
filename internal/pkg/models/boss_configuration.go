package models

import "time"

// Auth is a model of user credentials
type Auth struct {
	UseSSH bool   `json:"use,omitempty"`
	Path   string `json:"path,omitempty"`
	User   string `json:"x,omitempty"`
	Pass   string `json:"y,omitempty"`
}

// Configuration is a model of boss configurations
type Configuration struct {
	path                string
	Key                 string           `json:"id"`
	Auth                map[string]*Auth `json:"auth"`
	PurgeTime           int              `json:"purge_after"`
	InternalRefreshRate int              `json:"internal_refresh_rate"`
	LastPurge           time.Time        `json:"last_purge_cache"`
	LastInternalUpdate  time.Time        `json:"last_internal_update"`
	DelphiPath          string           `json:"delphi_path,omitempty"`
	ConfigVersion       int64            `json:"config_version"`
	GitEmbedded         bool             `json:"git_embedded"`
}
