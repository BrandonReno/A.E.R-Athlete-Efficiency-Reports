package data

import (
	"testing"
	"github.com/BrandonReno/A.E.R/data"
)

func TestValidation(t *testing.T){
	w := &data.Workout{
		User: data.Athlete{
			Name:  "Erik Clemensen",
			Sport: "rUnning",
		},
		Description: "Taper!",
	}

	err := w.Validate_Workout()

	if err != nil{
		t.Fatal(err)
	}
}