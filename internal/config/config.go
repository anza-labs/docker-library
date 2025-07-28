package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pelletier/go-toml/v2"
)

type Defaulter interface {
	Default()
}

func Load(file string, v Defaulter) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", file, err)
	}
	defer f.Close()

	if err := toml.NewDecoder(f).Decode(v); err != nil {
		return fmt.Errorf("failed to decode %s: %w", file, err)
	}

	v.Default()

	return nil
}

type Config struct {
	Use       string             `toml:"use"`
	BaseImage string             `toml:"base_image"`
	Labels    map[string]string  `toml:"labels"`
	Registry  *Registry          `toml:"registry"`
	Workflow  *Workflow          `toml:"workflow"`
	Engines   map[string]*Engine `toml:"engine"`
}

func (c *Config) Default() {
	c.Registry.Default()

	if c.BaseImage == "" {
		c.BaseImage = "ghcr.io/anza-labs/library/distroless/static-rootless:latest"
	}
}

type Registry struct {
	Name       string `toml:"name"`
	Username   string `toml:"username"`
	Password   string `toml:"password"`
	Repository string `toml:"repository"`
}

func (r *Registry) Default() {
	if r.Name == "" {
		r.Name = "ghcr.io"
	}

	if r.Repository == "" {
		r.Repository = os.Getenv("USER")
	}

	if r.Username == "" {
		r.Username = os.Getenv("USER")
	}

	if r.Password == "" {
		r.Password = "${{ secrets.GITHUB_TOKEN }}"
	}

}

type Workflow struct {
	Prefix string `toml:"prefix"`
}

func (w *Workflow) Default() {
	if w.Prefix == "" {
		w.Prefix = "gen_"
	}
}

type Engine struct {
	Build    []string `toml:"build"`
	Push     []string `toml:"push"`
	Manifest Manifest `toml:"manifest"`
}

type Manifest struct {
	Create   []string `toml:"create"`
	Annotate []string `toml:"annotate"`
	Push     []string `toml:"push"`
}

type Package struct {
	OS        string            `toml:"os"`
	BaseImage string            `toml:"base_image"`
	Arch      []string          `toml:"arch"`
	Versions  []string          `toml:"versions"`
	Targets   []string          `toml:"targets"`
	Labels    map[string]string `toml:"labels"`
}

func (p *Package) Default() {
	if p.OS == "" {
		p.OS = "linux"
	}

	if len(p.Arch) == 0 {
		p.Arch = []string{runtime.GOARCH}
	}

	if len(p.Versions) == 0 {
		p.Versions = []string{"latest"}
	}

	if len(p.Targets) == 0 {
		p.Targets = []string{"image"}
	}
}
