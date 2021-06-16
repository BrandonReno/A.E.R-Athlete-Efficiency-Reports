package pooling

type Job struct {
	Task func() error
}

func (j *Job) Process() error {
	return j.Task()
}