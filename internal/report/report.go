package report

type CoverageReport struct {
	Pipelines []PipelineCoverage `yaml:"pipelines"`
}

type PipelineCoverage struct {
	File string        `yaml:"file"`
	Jobs []JobCoverage `yaml:"jobs"`
}

type JobCoverage struct {
	Name    string           `yaml:"name"`
	Status  string           `yaml:"status"`
	Asserts []AssertCoverage `yaml:"asserts"`
}

type AssertCoverage struct {
	Name     string            `yaml:"name"`
	Result   string            `yaml:"result"`
	Artifact string            `yaml:"artifact,omitempty"`
	Checks   map[string]string `yaml:"checks,omitempty"` // key=string buscado, val=passed|failed
}
