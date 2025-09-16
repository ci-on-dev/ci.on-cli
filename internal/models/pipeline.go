package models

type Pipeline struct {
	Provider string         `yaml:"provider"`
	Stages   []string       `yaml:"stages"`
	File     string         `yaml:"file"`
	Vars     map[string]any `yaml:"vars"`
}
