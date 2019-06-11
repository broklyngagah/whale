package version

import (
	"fmt"
	"strings"
)

var (
	GitCommit string

	Version           string
	VersionPrerelease string
)

func GetHumanVersion() string {
	version := Version

	release := VersionPrerelease
	if release == "" {
		release = "dev"
	}
	if release != "" {
		version += fmt.Sprintf("-%s", release)
		if GitCommit != "" {
			version += fmt.Sprintf(" (%s)", GitCommit)
		}
	}

	return strings.Replace(version, "'", "", -1)
}
