package models

type Suite struct {
	Name     string   `yaml:"name"`
	Pipeline Pipeline `yaml:"pipeline"`
	Tests    []Test   `yaml:"tests"`
	Reports  []Report `yaml:"reports"`
}

type Test struct {
	Name    string         `yaml:"name"`
	Vars    map[string]any `yaml:"vars"`
	Asserts []Assert       `yaml:"asserts"`
	Stages  []string       `yaml:"stages"`
	After   []string       `yaml:"after"`
}

type RunCfg struct {
	Mode    string `yaml:"mode"`
	Network string `yaml:"network"`
}

type Artifact struct {
	File        string   `yaml:"file"`
	Type        string   `yaml:"type"`
	Contains    []string `yaml:"contains"`
	NotContains []string `yaml:"notContains"`
}

type Output struct {
	Contains []string `yaml:"contains"`
}

type Assert struct {
	Name   string `yaml:"name"`
	Expect Expect `yaml:"expect"`
}

type Expect struct {
	Job            string    `yaml:"job"`
	Status         string    `yaml:"status"`
	Step           string    `yaml:"step"`
	StepStatus     string    `yaml:"step_status"`
	Artifact       *Artifact `yaml:"artifact,omitempty"`
	OutputContains *Output   `yaml:"output,omitempty"`
}

type Report struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}
