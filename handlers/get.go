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

func GetSingleWorkout(rw http.ResponseWriter, r *http.Request){

	//swagger stuff here



	Athlete_ID := getAthleteID(r)
	Workout_ID, err := getWorkoutID(r)

	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in converting string to int: %s", err), http.StatusInternalServerError)
		return
	}

	w, err := data.GetSingleWorkout(Athlete_ID, Workout_ID)

	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in getting workout: %s", err), http.StatusBadRequest)
		return
	}

	err = data.ToJSON(w, rw)

	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in serializing workout: %s", err), http.StatusBadRequest)
	}
}


func GetAthlete(rw http.ResponseWriter, r *http.Request){

	//swagger here

	athlete_id := getAthleteID(r)
	athlete, err := data.GetAthlete(athlete_id)
	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in getting athlete: %s", err), http.StatusBadRequest)
		return
	}
	data.ToJSON(athlete, rw)
}