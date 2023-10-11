package job

// Interface ...
type Interface interface {
	Work() string
}

// job struct
type job struct {
}

// Work ...
func (j *job) Work() string {
	return "Work job"
}

// NewJob fabric for job with Interface
func NewJob() Interface {
	return &job{}
}
