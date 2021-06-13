package docs

// The ID number of the workout to be deleted
// swagger:parameters deleteWorkout updateWorkout getSingleWorkout
type WorkoutIDParameterWrapper struct {
	//	in: path
	//	required: True
	// 	example: 374638
	Workout_id int `json:"workout_id"`
}

// The unique athlete ID which distinguishes athletes
// swagger:parameters getSingleWorkout getAthlete deleteAthlete updateAthlete updateWorkout addWorkout deleteWorkout getSingleWorkout getEfficiency
type AthleteIDParameterWrapper struct {
	//	in:path
	//	required: true
	//	example:H3bfj78eHe
	Athlete_id string `json:"athlete_id"`
}




