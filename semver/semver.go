package semver

import (
	"fmt"
)

// Version hold the current repo version.
var (
	Version version = version{}

	Build = Version.Build
)

// version returns the parsed form of a semantic version string.
type version struct {
	major      string
	minor      string
	patch      string
	short      string
	prerelease string
	build      string
	err        string
}

// Build returns the build suffix of the semantic version v.
// For example, Build("v2.1.0+meta") == "+meta".
// If v is an invalid semantic version string, Build returns the empty string.
func (v version) Build() string {
	return v.build
}

func (v version) Set(s string) error {
	pv, ok := parse(s)
	if !ok {
		return fmt.Errorf("version format is invalid: %v", pv)
	}
	v.major = pv.major
	v.minor = pv.minor
	v.patch = pv.patch
	v.short = pv.short
	v.prerelease = pv.prerelease
	v.build = pv.build
	v.err = pv.err
	return nil
}

// IsValid reports whether v is a valid semantic version string.
func IsValid(v string) bool {
	_, ok := parse(v)
	return ok
}

// func GetVersionTag(v string) string {

// 	version := zsh.Version()
// 	fmt.Println("version: ", version)
// 	if version == "" || version[:5] == `/x1b[` || !semver.IsValid(version) {
// 		version = "unknown"
// 	}
// 	v := semver.MajorMinor(version)
// }

// func GetVersion() {}
