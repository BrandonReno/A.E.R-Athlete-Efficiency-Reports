package pooling

type Job struct {
	Name string //name used for logging to say what task is being processed
	Task func() error 
}

// return the result of running the task
func (j *Job) Process() error {
	return j.Task()
}