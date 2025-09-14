package builder

import (
	"docker-library/internal/config"
	"fmt"
	"path"
	"slices"
)

type Engine struct {
	Name             string
	Build            []string
	Push             []string
	ManifestCreate   []string
	ManifestAnnotate []string
	ManifestPush     []string
}

func (e *Engine) BuildCommand(
	registry, repository, name, os, arch, target, version, baseImage string,
	labels map[string]string,
	mono bool,
) []string {
	full := path.Join(registry, repository, name)
	if !mono {
		full = path.Join(full, target)
	}

	labelsMerged := []string{}
	for k, v := range labels {
		labelsMerged = append(labelsMerged, fmt.Sprintf("--label=%s='%s'", k, v))
	}

	slices.Sort(labelsMerged)

	if len(labelsMerged) > 0 {
		e.Build = append(
			e.Build,
			labelsMerged...,
		)
	}

	return slices.Concat(
		[]string{e.Name}, e.Build,
		[]string{
			fmt.Sprintf("--target=%s", target),
			fmt.Sprintf("--build-arg=VERSION=%s", version),
			fmt.Sprintf("--build-arg=BASE_IMAGE=%s", baseImage),
			fmt.Sprintf("--platform=%s/%s", os, arch),
			fmt.Sprintf("--file=%s", path.Join("./", name, "Dockerfile")),
			fmt.Sprintf("--tag=%s:%s-%s", full, version, arch),
			path.Join("./", name),
		},
	)
}

func (e *Engine) PushCommand(registry, repository, name, os, arch, target, version string, mono bool) []string {
	full := path.Join(registry, repository, name)
	if !mono {
		full = path.Join(full, target)
	}
	return slices.Concat(
		[]string{e.Name}, e.Push,
		[]string{fmt.Sprintf("%s:%s-%s", full, version, arch)},
	)
}

func (e *Engine) ManifestCreateCommand(registry, repository, name, target, version string, mono bool, pkg *config.Package) []string {
	full := path.Join(registry, repository, name)
	if !mono {
		full = path.Join(full, target)
	}

	archImages := []string{}

	for _, arch := range pkg.Arch {
		archImages = append(archImages, fmt.Sprintf("%s:%s-%s", full, version, arch))
	}

	return slices.Concat(
		[]string{e.Name}, e.ManifestCreate,
		[]string{fmt.Sprintf("%s:%s", full, version)},
		archImages,
	)
}

func (e *Engine) ManifestAnnotateCommand(registry, repository, name, os, arch, target, version string, mono bool) []string {
	full := path.Join(registry, repository, name)
	if !mono {
		full = path.Join(full, target)
	}

	return slices.Concat(
		[]string{e.Name}, e.ManifestAnnotate,
		[]string{
			fmt.Sprintf("--arch=%s", arch),
			fmt.Sprintf("%s:%s", full, version),
			fmt.Sprintf("%s:%s-%s", full, version, arch),
		},
	)
}

func (e *Engine) ManifestPushCommand(registry, repository, name, target, version string, mono bool) []string {
	full := path.Join(registry, repository, name)
	if !mono {
		full = path.Join(full, target)
	}
	return slices.Concat(
		[]string{e.Name}, e.ManifestPush,
		[]string{fmt.Sprintf("%s:%s", full, version)},
	)
}
