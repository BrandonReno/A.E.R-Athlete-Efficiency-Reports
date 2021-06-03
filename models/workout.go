package models

import (
	"strings"
	"github.com/go-playground/validator")


type Workout struct {
	Workout_ID  int    `json:"workout_id"`
	Athlete_ID  string `json:"-" validate:"required"`
	Date        string `json:"date"`
	Description string `json:"description" validate:"required"`
	Sport       string `json:"sport" validate:"Sport"`
	Rating      int    `json:"rating" validate:"gte=0,lte=10"`
}

var AvailableSports = []string{"swimming", "running", "lifting", "biking"}

func (w *Workout) Validate_Workout() error {
	validate := validator.New()                              //Create a new validator and hold it in var validate
	validate.RegisterValidation("Sport", validateSport) //register the Sport field with the validate sport function
	err := validate.Struct(w)                           //validate the struct w and return any errors
	return err
}

func validateSport(fl validator.FieldLevel) bool {
	sport := fl.Field().String()         //get the string value of the sport field
	for _, sp := range AvailableSports { //iterate through accessible sports
		if strings.ToLower(sport) == strings.ToLower(sp) { // if the field sport is in the acceptable sports return true else false
			return true
		}
	}
	return false
}