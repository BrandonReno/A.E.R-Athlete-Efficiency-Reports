package docs

import "github.com/BrandonReno/A.E.R/models"

//A list of workouts returns in the response
// swagger:response workoutsResponse
type workoutsResponseWrapper struct{
	//	in:body
		body []models.Workout
}

//A single workout returns in the response
// swagger:response singleWorkout
type singleWorkoutResponseWrapper struct{
	// in: body
	Body models.Workout
}


// No content is returned
// swagger:response noContent
type WorkoutNoContentWrapper struct{
}

// Index Not Found
// swagger:response badRequest
type badRequestWrapper struct{
	// 	example: Could not delete/update workout
		err string `json:"error"`
	}
	
// Validation Error
// swagger:response verror
type validationErrorWrapper struct{
//	example: Unable to validate JSON object
	err string `json:"error"`
}

