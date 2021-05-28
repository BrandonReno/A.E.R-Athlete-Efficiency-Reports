package controllers

import (
	"fmt"
	"net/http"
	"github.com/BrandonReno/A.E.R/services"
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

	wl, err := services.GetUserWorkouts(id)
	if err != nil{
		http.Error(rw, fmt.Sprintf("error returned: %s", err), http.StatusInternalServerError)
	}
	err = ToJSON(wl,rw) //Encode the list from structs to JSON objects
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

	w, err := services.GetSingleWorkout(Athlete_ID, Workout_ID)

	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in getting workout: %s", err), http.StatusBadRequest)
		return
	}

	err = ToJSON(w, rw)

	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in serializing workout: %s", err), http.StatusBadRequest)
	}
}


func GetAthlete(rw http.ResponseWriter, r *http.Request){

	//swagger here

	athlete_id := getAthleteID(r)
	athlete, err := services.GetAthlete(athlete_id)
	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in getting athlete: %s", err), http.StatusBadRequest)
		return
	}
	ToJSON(athlete, rw)
}