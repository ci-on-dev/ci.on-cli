package ir

type IR struct {
	Jobs       []Job
	Vars       map[string]any
	SourceFile string // <- aqui
}

type Job struct {
	Name  string
	Steps []Step
}

type Step struct {
	Name   string
	Run    string // comando shell
	Status string // "success" ou "failed"
	Logs   []string
}

// FindJob retorna ponteiro para job ou nil se não existir
func (ir *IR) FindJob(name string) *Job {
	for i := range ir.Jobs {
		if ir.Jobs[i].Name == name {
			return &ir.Jobs[i]
		}
	}
	return nil
}

// HasStep retorna true se job possui step com o nome
func (j *Job) HasStep(name string) bool {
	for _, s := range j.Steps {
		if s.Name == name {
			return true
		}
	}
	return false
}
