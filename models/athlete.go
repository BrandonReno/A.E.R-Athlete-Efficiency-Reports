package models

import "github.com/go-playground/validator"

//Athlete Structure which defines the 'User' of a workout
type Athlete struct {
	Athlete_ID string    `json:"athlete_id"`
	First_Name string    `json:"first_name" validate:"required"`
	Last_Name  string    `json:"last_name" validate:"required"`
	Age        uint8     `json:"age" validate:"gte=18, lte=100`
	Joined     string `json:"joined"`
}

func (a *Athlete) Validate_Athlete() error {
	validate := validator.New()
	err := validate.Struct(a)
	return err
}
