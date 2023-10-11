package activities

// job struct
type job struct {
}

// Do ...
func (j *job) Do() string {
	return "Work job"
}

// NewJob fabric for job with Activity interface
func NewJob() Activity {
	return &job{}
}
