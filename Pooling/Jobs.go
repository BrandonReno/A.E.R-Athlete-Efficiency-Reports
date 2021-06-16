package pooling

type Job struct{
	Data interface{}
	Task func(interface{})
}

func (j *Job) Process(){
	j.Task(j.Data)
}