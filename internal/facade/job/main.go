package job

type Interface interface {
	Work() string
}

type job struct {
}

func (j *job) Work() string {
	return "Work job"
}

func NewJob() Interface {
	return &job{}
}
