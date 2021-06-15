package pooling

import (
	"github.com/BrandonReno/A.E.R/services"
)

type Job struct{
	Name string
	Data interface{}
	Task func()
	Result error
}


func (j *Job) New(name string, data *interface{}, task func(), request string) *Job{
	return &Job{
		Name: name,
		Data: data,
		Task: task,
		Result: nil,
	}
}

func (j *Job) Process(){
	j.Task()
}