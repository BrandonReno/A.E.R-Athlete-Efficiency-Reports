package data

import "testing"

func TestValidation(t *testing.T){
	w := &Workout{
		User: Athlete{
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