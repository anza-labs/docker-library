//go:build mage
// +build mage

package main

import (
	"context"
	"docker-library/internal/actions"
	"docker-library/internal/builder"
	"docker-library/internal/config"
	"fmt"
	"os"
	"os/exec"
	"path"

	"gopkg.in/yaml.v3"
)

var (
	cfg = &config.Config{}
)

func init() {
	if err := config.Load("build.toml", cfg); err != nil {
		panic(err)
	}
}

func write(name string, content any) error {
	f, err := os.OpenFile(path.Join(".github/workflows", name), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.SetIndent(2)

	if err := enc.Encode(content); err != nil {
		return err
	}

	return nil
}

// Build...
func Build(ctx context.Context, image, version, arch, target string) error {
	pkg := &config.Package{}

	if err := config.Load(path.Join(image, "pkg.toml"), pkg); err != nil {
		return err
	}

	use := cfg.Use
	if engine, ok := os.LookupEnv("ENGINE"); ok && engine != "" {
		use = engine
	}

	eng := &builder.Engine{
		Name:  use,
		Build: cfg.Engines[use].Build,
	}

	cmd := eng.BuildCommand(
		cfg.Registry.Name, cfg.Registry.Repository, image,
		"linux", arch, target, version, true,
	)

	handle := exec.CommandContext(ctx, cmd[0], cmd[1:]...)
	handle.Stderr = os.Stderr
	handle.Stdout = os.Stdout

	return handle.Run()
}

// Push...
func Push(ctx context.Context, image, version, arch, target string) error {
	pkg := &config.Package{}

	if err := config.Load(path.Join(image, "pkg.toml"), pkg); err != nil {
		return err
	}

	use := cfg.Use
	if engine, ok := os.LookupEnv("ENGINE"); ok && engine != "" {
		use = engine
	}

	eng := &builder.Engine{
		Name: use,
		Push: cfg.Engines[use].Push,
	}

	cmd := eng.PushCommand(
		cfg.Registry.Name, cfg.Registry.Repository, image,
		"linux", arch, target, version, true,
	)

	handle := exec.CommandContext(ctx, cmd[0], cmd[1:]...)
	handle.Stderr = os.Stderr
	handle.Stdout = os.Stdout

	return handle.Run()
}

// Actions...
func Actions(ctx context.Context) error {
	libdir := "library"

	de, err := os.ReadDir(libdir)
	if err != nil {
		return err
	}

	for _, entry := range de {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if !entry.IsDir() {
			continue
		}

		pkg := &config.Package{}

		name := path.Join(libdir, entry.Name())

		if err := config.Load(path.Join(name, "pkg.toml"), pkg); err != nil {
			return err
		}

		bld, err := actions.Release(name, cfg, pkg)
		if err != nil {
			return err
		}

		workflow := fmt.Sprintf("%s_%s_%s.yaml", cfg.Workflow.Prefix, libdir, entry.Name())

		if err := write(workflow, bld); err != nil {
			return err
		}

	}

	return nil
}
