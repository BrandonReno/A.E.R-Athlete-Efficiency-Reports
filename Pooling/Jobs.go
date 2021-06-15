package pooling

type Job struct{
	Name string
	Data interface{}
	Task func() error
	Result error
}


func (j *Job) New(name string, data *interface{}, task func() error, request string) *Job{
	return &Job{
		Name: name,
		Data: data,
		Task: task,
		Result: nil,
	}
}

func (j *Job) Process() error{
	return j.Task()
}