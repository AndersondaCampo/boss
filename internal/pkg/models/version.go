package models

import (
	"flag"
	"runtime"
)

var (
	version      = "v3.0.0"
	metadata     = ""
	gitCommit    = ""
	gitTreeState = ""
)

// BuildInfo basic structure with compilation data
type BuildInfo struct {
	// Version is the current semver.
	Version string `json:"version,omitempty"`
	// GitCommit is the git sha1.
	GitCommit string `json:"git_commit,omitempty"`
	// GitTreeState is the state of the git tree.
	GitTreeState string `json:"git_tree_state,omitempty"`
	// GoVersion is the version of the Go compiler used.
	GoVersion string `json:"go_version,omitempty"`
}

func getVersion() string {
	if metadata == "" {
		return version
	}
	return version + "+" + metadata
}

// MakeBossVersion return boss version info
func MakeBossVersion() BuildInfo {
	v := BuildInfo{
		Version:      getVersion(),
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		GoVersion:    runtime.Version(),
	}
	if flag.Lookup("test.v") != nil {
		v.GoVersion = ""
	}
	return v
}
