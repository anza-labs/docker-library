package actions

import (
	_ "embed"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed workflow-templates/template.yaml
var actionsYaml string

var actions = make(map[string]string)

func loadActionsFromTemplate() error {
	var data struct {
		Jobs map[string]struct {
			Steps []struct {
				ID   string `yaml:"id"`
				Uses string `yaml:"uses"`
			} `yaml:"steps"`
		} `yaml:"jobs"`
	}

	decoder := yaml.NewDecoder(strings.NewReader(actionsYaml))
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	buildJob, ok := data.Jobs["build"]
	if !ok {
		return nil // or return error if build job is required
	}

	for _, step := range buildJob.Steps {
		if step.ID != "" && step.Uses != "" {
			actions[step.ID] = step.Uses
		}
	}
	return nil
}

func init() {
	if err := loadActionsFromTemplate(); err != nil {
		panic(err)
	}
}
