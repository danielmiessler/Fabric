package util

import (
	"runtime/debug"
	"strings"
)

// FabricVersion returns the module version Fabric was built with, or "dev"
// when it cannot be determined (for example, local builds without version
// information). Build metadata such as the "+dirty" VCS marker is stripped so
// the value is safe to use in a User-Agent.
func FabricVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		if version := normalizeVersion(info.Main.Version); version != "" {
			return version
		}
	}
	return "dev"
}

func normalizeVersion(version string) string {
	version = strings.TrimSpace(version)
	if version == "" || version == "(devel)" {
		return ""
	}
	if idx := strings.IndexByte(version, '+'); idx >= 0 {
		version = strings.TrimSpace(version[:idx])
	}
	return version
}
