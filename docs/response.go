package docs

import "github.com/BrandonReno/A.E.R/models"

//A list of workouts returns in the response
// swagger:response workoutsResponse
type workoutsResponseWrapper struct{
	//	in:body
	Body []models.Workout
}

//A single workout returns in the response
// swagger:response singleWorkout
type singleWorkoutResponseWrapper struct{
	// in: body
	Body models.Workout
}

// A single athlete returns in the response
// swagger:response singleAthlete
type singleAthleteResponseWrapper struct{
	// in:body
	Body models.Athlete
}

// All athletes are returned in the response
// swagger:response athletesResponse
type athletesResponseWrapper struct{
	// in:body
	Body []models.Athlete
}

// An athletes aer is returned in the response
// swagger:response efficiencyResponse
type ResponseEfficiencyWrapper struct{
	//	in:body
	Body models.Efficiency	
}


// No content is returned
// swagger:response noContent
type NoContentWrapper struct{
}

// Index Not Found
// swagger:response badRequest
type badRequestWrapper struct{
// 	example: Could not find the workout or athlete ID in the database
	Err string `json:"error"`
}
	
// Validation Error
// swagger:response verror
type validationErrorWrapper struct{
//	example: Unable to validate athlete or workout object
	Err string `json:"error"`
}

