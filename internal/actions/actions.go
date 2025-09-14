package actions

import (
	"docker-library/internal/builder"
	"docker-library/internal/config"
	"fmt"
	"maps"
	"path"
	"strings"
)

func buildMatrix(pkg *config.Package) map[string]any {
	return map[string]any{
		"version": pkg.Versions,
		"arch":    pkg.Arch,
		"target":  pkg.Targets,
	}
}

func promoteMatrix(pkg *config.Package) map[string]any {
	return map[string]any{
		"version": pkg.Versions,
		"target":  pkg.Targets,
	}
}

func runManifestCreate(name string, cfg *config.Config, pkg *config.Package) string {
	engine := &builder.Engine{
		Name:             cfg.Use,
		ManifestCreate:   cfg.Engines[cfg.Use].Manifest.Create,
		ManifestAnnotate: cfg.Engines[cfg.Use].Manifest.Annotate,
		ManifestPush:     cfg.Engines[cfg.Use].Manifest.Push,
	}

	emc := engine.ManifestCreateCommand(
		cfg.Registry.Name, cfg.Registry.Repository, name,
		"${{ matrix.target }}", "${{ matrix.version }}",
		len(pkg.Targets) == 1, pkg,
	)

	return strings.Join(emc, " \\\n    ")
}

func runManifestAnnotate(name string, cfg *config.Config, pkg *config.Package) string {
	engine := &builder.Engine{
		Name:             cfg.Use,
		ManifestCreate:   cfg.Engines[cfg.Use].Manifest.Create,
		ManifestAnnotate: cfg.Engines[cfg.Use].Manifest.Annotate,
		ManifestPush:     cfg.Engines[cfg.Use].Manifest.Push,
	}

	script := []string{}

	for _, arch := range pkg.Arch {
		ema := engine.ManifestAnnotateCommand(
			cfg.Registry.Name, cfg.Registry.Repository, name, pkg.OS, arch,
			"${{ matrix.target }}", "${{ matrix.version }}",
			len(pkg.Targets) == 1,
		)
		script = append(script, strings.Join(ema, " \\\n    "))
	}

	return strings.Join(script, " && \\\n")
}

func runManifestPush(name string, cfg *config.Config, pkg *config.Package) string {
	engine := &builder.Engine{
		Name:             cfg.Use,
		ManifestCreate:   cfg.Engines[cfg.Use].Manifest.Create,
		ManifestAnnotate: cfg.Engines[cfg.Use].Manifest.Annotate,
		ManifestPush:     cfg.Engines[cfg.Use].Manifest.Push,
	}

	emp := engine.ManifestPushCommand(
		cfg.Registry.Name, cfg.Registry.Repository, name,
		"${{ matrix.target }}", "${{ matrix.version }}",
		len(pkg.Targets) == 1,
	)

	return strings.Join(emp, " \\\n    ")
}

func runBuild(name string, cfg *config.Config, pkg *config.Package) string {
	engine := &builder.Engine{
		Name:  cfg.Use,
		Build: cfg.Engines[cfg.Use].Build,
	}

	labels := maps.Clone(cfg.Labels)
	maps.Copy(labels, pkg.Labels)

	baseImage := pkg.BaseImage
	if baseImage == "" {
		baseImage = cfg.BaseImage
	}

	return strings.Join(engine.BuildCommand(
		cfg.Registry.Name, cfg.Registry.Repository, name, pkg.OS,
		"${{ matrix.arch }}", "${{ matrix.target }}", "${{ matrix.version }}",
		baseImage,
		labels, len(pkg.Targets) == 1,
	), " \\\n    ")
}

func runPush(name string, cfg *config.Config, pkg *config.Package) string {
	engine := &builder.Engine{
		Name: cfg.Use,
		Push: cfg.Engines[cfg.Use].Push,
	}

	return strings.Join(engine.PushCommand(
		cfg.Registry.Name, cfg.Registry.Repository, name, pkg.OS,
		"${{ matrix.arch }}", "${{ matrix.target }}", "${{ matrix.version }}",
		len(pkg.Targets) == 1,
	), " \\\n    ")
}

func runSign(name string, cfg *config.Config, pkg *config.Package) string {
	cosign := &builder.Cosign{}

	return strings.Join(cosign.SignImageCommand(
		cfg.Registry.Name, cfg.Registry.Repository, name,
		"${{ matrix.arch }}", "${{ matrix.target }}", "${{ matrix.version }}",
		len(pkg.Targets) == 1,
	), " \\\n    ")
}

func runManifestSign(name string, cfg *config.Config, pkg *config.Package) string {
	cosign := &builder.Cosign{}

	return strings.Join(cosign.SignImageCommand(
		cfg.Registry.Name, cfg.Registry.Repository, name,
		"", "${{ matrix.target }}", "${{ matrix.version }}",
		len(pkg.Targets) == 1,
	), " \\\n    ")
}

func Release(name string, cfg *config.Config, pkg *config.Package) (map[string]any, error) {
	return map[string]any{
		"name": fmt.Sprintf("release %s", name),
		"on": map[string]any{
			"push": map[string]any{
				"branches": []string{"main"},
				"paths": []string{
					path.Join(".github/workflows",
						fmt.Sprintf("%s_%s.yaml", cfg.Workflow.Prefix, strings.ReplaceAll(name, "/", "_"))),
					path.Join(name, "Dockerfile"),
					path.Join(name, "pkg.toml"),
				},
			},
			"workflow_dispatch": struct{}{},
		},
		"permissions": map[string]string{
			"contents": "write",
			"packages": "write",
			"id-token": "write",
		},
		"jobs": map[string]any{
			"build": map[string]any{
				"runs-on": "ubuntu-latest",
				"env":     map[string]string{},
				"strategy": map[string]any{
					"matrix": buildMatrix(pkg),
				},
				"steps": []map[string]any{
					{
						"uses": actions["actionCheckout"],
					},
					{
						"uses": actions["actionDockerLogin"],
						"with": map[string]any{
							"registry": cfg.Registry.Name,
							"username": cfg.Registry.Username,
							"password": cfg.Registry.Password,
						},
					},
					{
						"uses": actions["actionSetupQEMU"],
					},
					{
						"uses": actions["actionSetupBuildx"],
					},
					{
						"run": runBuild(name, cfg, pkg),
					},
					{
						"run": runPush(name, cfg, pkg),
					},
					{
						"uses": actions["actionCosignInstaller"],
					},
					{
						"run": runSign(name, cfg, pkg),
					},
				},
			},
			"promote": map[string]any{
				"runs-on": "ubuntu-latest",
				"needs":   "build",
				"env":     map[string]string{},
				"strategy": map[string]any{
					"matrix": promoteMatrix(pkg),
				},
				"steps": []map[string]any{
					{
						"uses": actions["actionCheckout"],
					},
					{
						"uses": actions["actionDockerLogin"],
						"with": map[string]any{
							"registry": cfg.Registry.Name,
							"username": cfg.Registry.Username,
							"password": cfg.Registry.Password,
						},
					},
					{
						"uses": actions["actionSetupBuildx"],
					},
					{
						"run": runManifestCreate(name, cfg, pkg),
					},
					{
						"run": runManifestAnnotate(name, cfg, pkg),
					},
					{
						"run": runManifestPush(name, cfg, pkg),
					},
					{
						"uses": actions["actionCosignInstaller"],
					},
					{
						"run": runManifestSign(name, cfg, pkg),
					},
				},
			},
		},
	}, nil
}

func Scan(name string, cfg *config.Config, pkg *config.Package) (map[string]any, error) {
	return map[string]any{
		"name": fmt.Sprintf("scan %s", name),
	}, nil
}
