package docs

import "github.com/BrandonReno/A.E.R/models"

// The ID number of the workout to be deleted
// swagger:parameters deleteWorkout updateWorkout getWorkout
type WorkoutIDParameterWrapper struct{
	//	in: path
	//	required: True
	// 	example: 374638
		ID int `json:"id"`
	}

// The Workout parameters
// swagger:parameters addWorkout updateWorkout
type RequestWorkoutWrapper struct{
	//	in: body
	//	required: True
	Body []models.Workout
}