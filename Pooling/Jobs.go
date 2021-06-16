package pooling

type Job struct{
	Data interface{}
	Task func(interface{}) error
}

func (j *Job) Process()error{
	return j.Task(j.Data)
}