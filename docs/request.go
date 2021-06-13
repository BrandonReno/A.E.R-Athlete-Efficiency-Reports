package docs

import "github.com/BrandonReno/A.E.R/models"

// The Workout parameters
// swagger:parameters addWorkout updateWorkout
type RequestWorkoutWrapper struct {
	//	in: body
	//	required: True
	Body models.Workout
}

// The parameters to be filled out when creating an athlete
// swagger:parameters addAthlete updateAthlete
type RequestAthleteWrapper struct {
	//	in: body
	//	required: true
	Body models.Athlete
}