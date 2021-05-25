package handlers

import (
	"fmt"
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
	id := getAthleteID(r)

	wl, err := data.GetUserWorkouts(id)
	if err != nil{
		http.Error(rw, fmt.Sprintf("error returned: %s", err), http.StatusInternalServerError)
	}
	err = data.ToJSON(wl,rw) //Encode the list from structs to JSON objects
	if err != nil{ //if json can not be encoded return an error and log the error while also returning out of the function
		http.Error(rw, "Unable to encode JSON object", http.StatusInternalServerError)
		return 
	}
}
