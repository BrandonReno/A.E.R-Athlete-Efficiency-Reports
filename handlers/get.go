package handlers

import (
	"net/http"
	"github.com/BrandonReno/A.E.R/data"
)

// Gets all workouts in the system
func GetWorkouts(rw http.ResponseWriter, r *http.Request){

	// swagger:route GET / workouts listWorkouts
    //
    // Lists workouts from all athletes in the system
    //
    //     Produces:
    //     - application/json
    //
    //     Schemes: http
    //
    //     Responses: 
	//			200: workoutsResponse
	
	wl := data.GetFeed() //Get the list of workouts from 'data'
	err := data.ToJSON(wl,rw) //Encode the list from structs to JSON objects
	if err != nil{ //if json can not be encoded return an error and log the error while also returning out of the function
		http.Error(rw, "Unable to encode JSON object", http.StatusInternalServerError)
		return 
	}
}


func GetSingleWorkout(rw http.ResponseWriter, r *http.Request){
	// swagger:route GET /workouts/wid{id} workouts getWorkout
	//
	// Displays a single workout based upon the inputted ID
	//
	//		Produces:
	//		- application/json
	//
	//		Schemes: http
	//
	//		Responses:
	//			200: singleWorkout
	//			400: verror

	id := getProductID(r)

	w, _, err := data.FindWorkout(id)

	if err != nil{
		http.Error(rw, "Could not find Workout with specified ID", http.StatusNotFound)
		return
	}

	err = data.ToJSON(w, rw)

	if err != nil{
		http.Error(rw, "Could not serialize workout object", http.StatusInternalServerError)
		return
	}
}