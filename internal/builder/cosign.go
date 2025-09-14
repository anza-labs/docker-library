package builder

import (
	"fmt"
	"path"
)

type Cosign struct{}

func (c *Cosign) SignImageCommand(registry, repository, name, arch, target, version string, mono bool) []string {
	full := path.Join(registry, repository, name)
	if !mono {
		full = path.Join(full, target)
	}

	tag := version
	if arch != "" {
		tag = fmt.Sprintf("%s-%s", tag, arch)
	}

	return []string{
		"cosign",
		"sign",
		"--keyless",
		fmt.Sprintf("%s:%s", full, tag),
	}
}
