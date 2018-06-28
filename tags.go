package docker

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

// DefaultTags returns a set of default suggested tags based on
// the commit ref.
func DefaultTags(ref string) []string {
	if strings.HasPrefix(ref, "refs/heads/") {
		branch := stripHeadPrefix(ref)
		tag := branch
		if tag == "master" {
			tag = "latest"
		}
		return []string{tag}
	}
	if !strings.HasPrefix(ref, "refs/tags/") {
		return []string{"dev"}
	}
	v := stripTagPrefix(ref)
	version, err := semver.NewVersion(v)
	if err != nil {
		return []string{"dev"}
	}
	if version.PreRelease != "" || version.Metadata != "" {
		return []string{
			version.String(),
		}
	}
	if version.Major == 0 {
		return []string{
			"latest",
			fmt.Sprintf("%d.%d", version.Major, version.Minor),
			fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch),
		}
	}
	return []string{
		"latest",
		fmt.Sprint(version.Major),
		fmt.Sprintf("%d.%d", version.Major, version.Minor),
		fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch),
	}
}

func stripHeadPrefix(ref string) string {
	return strings.TrimPrefix(ref, "refs/heads/")
}

func stripTagPrefix(ref string) string {
	ref = strings.TrimPrefix(ref, "refs/tags/")
	ref = strings.TrimPrefix(ref, "v")
	return ref
}
