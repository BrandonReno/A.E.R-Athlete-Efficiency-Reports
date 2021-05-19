package handlers

import (
	"net/http"
	"github.com/BrandonReno/A.E.R-Athlete-Efficiency-Reports/data"
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
	err := wl.ToJSON(rw) //Encode the list from structs to JSON objects
	if err != nil{ //if json can not be encoded return an error and log the error while also returning out of the function
		http.Error(rw, "Unable to encode JSON object", http.StatusInternalServerError)
		return 
	}
}