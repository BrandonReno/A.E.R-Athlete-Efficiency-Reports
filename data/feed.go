package data

import (
	"encoding/json"
	"fmt"
	"io"
	"github.com/go-playground/validator"
	"strings"
)

//Workout Structure which holds details about an athletes workout
type Workout struct { 
	ID          int     `json:"id"`
	User        Athlete `json:"user" validate:"required"`
	Date        string  `json:"date"`
	Description string  `json:"description" validate:"required"`
}

//Athlete Structure which defines the 'User' of a workout
type Athlete struct {
	ID    int    `json:"athlete_id"`
	Name  string `json:"name" validate:"required"`
	Sport string `json:"sport" validate:"required,Sport"`
	Age   uint8    `json:"age" validate:"gte=18, lte=100`
}

var AvailableSports = []string{"swimming", "running", "lifting", "biking",}

var ErrorWorkoutNotFound = fmt.Errorf("Workout id not found")

type Workout_Feed []*Workout //Create Workout_Feed of slice of referenced Workout

func (w *Workout) Validate_Workout() error{
	validate := validator.New() //Create a new validator and hold it in var validate
	validate.RegisterValidation("Sport", validateSport) //register the Sport field with the validate sport function
	err := validate.Struct(w) //validate the struct w and return any errors
	return err
}

func validateSport(fl validator.FieldLevel) bool{
	sport := fl.Field().String() //get the string value of the sport field
	for _, sp := range AvailableSports{ //iterate through accessible sports
		if strings.ToLower(sport) == strings.ToLower(sp){ // if the field sport is in the acceptable sports return true else false
			return true
		}
	}
	return false
}

func (wf *Workout_Feed) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(wf) // Create a new encoder and encode the current Workout_Feed to json. Returns an error just in case
}

func (w *Workout) FromJSON(r io.Reader) error{
	return json.NewDecoder(r).Decode(w) // Create a new decoder and decode the request body to json. Returns an error just in case
}

//Definition of private workoutFeed list
var workoutFeed = []*Workout{
	&Workout{
		ID: 1,
		User: Athlete{
			ID:    1,
			Name:  "Brandon Reno",
			Sport: "Swimming",
			Age:   23,
		},
		Date:        "5/13/2021",
		Description: "Hard Workout",
	},

	&Workout{
		ID: 2,
		User: Athlete{
			ID:    2,
			Name:  "Kelsey LLoyd",
			Sport: "Swimming",
			Age:   20,
		},
		Date:        "5/11/2021",
		Description: "Tough!",
	},

	&Workout{
		ID: 3,
		User: Athlete{
			ID:    3,
			Name:  "Erik Clemensen",
			Sport: "Swimming",
			Age:   22,
		},
		Date:        "5/13/2021",
		Description: "Taper!",
	},
}
