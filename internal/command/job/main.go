package job

import "main/internal/command/models"

type job struct {
}

func (j *job) Do() string {
	return "Work job"
}

func NewJob() models.Activity {
	return &job{}
}
