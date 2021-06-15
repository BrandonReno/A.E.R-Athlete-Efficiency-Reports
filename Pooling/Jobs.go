package pooling

import (
	"github.com/BrandonReno/A.E.R/services"
)

type Job struct{
	Name string
	Data interface{}
	Database *services.DB
	Result error
}


func (j *Job) New(name string, data *interface{}, database *services.DB) *Job{
	return &Job{
		Name: name,
		Data: data,
		Database: database,
		Result: nil,
	}
}

func (j *Job) Process(){
}