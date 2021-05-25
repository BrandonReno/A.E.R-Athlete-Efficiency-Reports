package data

import(
	"github.com/go-playground/validator"
	"strings"
)

type Efficiency struct {
	Efficiency_ID    int     `json:"efficiency_id"`
	Efficiency_Score float64 `json:"efficiency_score"`
	Favorite_Sport   string  `json:"favorite_excercise"`
	Athlete_ID       string  `json:"athlete_id"`
}

//Workout Structure which holds details about an athletes workout
type Workout struct {
	Workout_ID  int    `json:"-"`
	Athlete_ID  string `json:"-" validate:"required"`
	Date        string `json:"date"`
	Description string `json:"description" validate:"required"`
	Sport       string `json:"sport" validate Sport`
	Rating      int    `json:"rating" validate: gte=0, lte=10`
}

//Athlete Structure which defines the 'User' of a workout
type Athlete struct {
	Athlete_ID int    `json:"athlete_id"`
	First_Name string `json:"first_name" validate:"required"`
	Last_Name  string `json:"last_name" validate:"required"`
	Age        uint8  `json:"age" validate:"gte=18, lte=100`
	Joined     string `json:"joined"`
}


type EfficiencyViewWrapper struct {
	First_Name string `json:"first_name" validate:"required"`
	Last_Name  string `json:"last_name" validate:"required"`
	Favorite_Sport   string  `json:"favorite_excercise"`
	Efficiency_Score float64 `json:"efficiency_score"`
}


var validate *validator.Validate

var AvailableSports = []string{"swimming", "running", "lifting", "biking",}

func (w *Workout) Validate_Workout() error{
	validate = validator.New() //Create a new validator and hold it in var validate
	validate.RegisterValidation("Sport", validateSport) //register the Sport field with the validate sport function
	err := validate.Struct(w) //validate the struct w and return any errors
	return err
}

func (a *Athlete) Validate_Athlete() error{
	validate = validator.New()
	err := validate.Struct(a)
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