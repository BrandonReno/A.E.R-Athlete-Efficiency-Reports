// Package Classification AER: Workout API
//
// The purpose of this API is to handle back end
// AER services following RESTful principles
//
//	Schemes: http
//  BasePath: /
//	Version: 1.0.0
//	
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	swagger:meta
package docs

import "github.com/BrandonReno/Workouts/data"

//A list of workouts returns in the response
// swagger:response workoutsResponse
type workoutsResponseWrapper struct{
//	in:body
	body []data.Workout
}

// The ID number of the workout to be deleted
// swagger:parameters deleteWorkout updateWorkout
type WorkoutIDParameterWrapper struct{
//	in: path
//	required: True
// 	example: 374638
	ID int `json:"id"`
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

// swagger:model
type Workout struct { 
	// The id for this workout
    //
    // min: 1
	// read only: true
	ID          int     `json:"id"`
	// The user whom completed this workout
    //
    // required: true
	User        Athlete `json:"user" validate:"required"`
	// The date this workout was completed
	//
	// read only: true
	Date        string  `json:"date"`
	// The description of the workout
    //
    // maximum length: 500
	// required: true
	Description string  `json:"description" validate:"required"`
}

//Athlete Structure which defines the 'User' of a workout
type Athlete struct {
	// The ID linking to the user
    //
    // min: 1
	// Read Only: true
	ID    int    `json:"athlete_id"`
	// The full name of the user
    //
    // minimum length: 5
	// maximum length: 50
	// required: true
	Name  string `json:"name" validate:"required"`
	// The sport the user does
    //
	// required: true
	Sport string `json:"sport" validate:"required,Sport"`
	// The age of the user
    //
    // min: 18
	// max: 100
	Age   uint8    `json:"age" validate:"gte=18, lte=100`
}

