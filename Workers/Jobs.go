package workers

import (
	"github.com/BrandonReno/A.E.R/services"
)

type Job struct{
	Name string
	Data interface{}
	Database *services.DB
	Result error
}