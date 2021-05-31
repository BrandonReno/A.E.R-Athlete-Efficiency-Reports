package docs

import "github.com/BrandonReno/A.E.R/models"

// The ID number of the workout to be deleted
// swagger:parameters deleteWorkout updateWorkout getWorkout
type WorkoutIDParameterWrapper struct {
	//	in: path
	//	required: True
	// 	example: 374638
	ID int `json:"workout_id"`
}

// The unique athlete ID which distinguishes athletes
// swagger:parameters getSingleWorkout getAthlete deleteAthlete updateAthlete updateWorkout
type AthleteIDParameterWrapper struct {
	//	in:path
	//	required: true
	//	example:H3bfj78eHe
	ID string `json:"athlete_id"`
}

// The Workout parameters
// swagger:parameters addWorkout updateWorkout
type RequestWorkoutWrapper struct {
	//	in: body
	//	required: True
	Body models.Workout
}

// The parameters to be filled out when creating an athlete
// swagger:parameters addAthlete
type RequestAthleteWrapper struct {
	//	in: body
	//	required: true
	Body models.Athlete
}
