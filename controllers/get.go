package controllers

import (
	"fmt"
	"net/http"
)

// Gets all workouts in the system
func (l *Aer_Log) GetWorkouts(rw http.ResponseWriter, r *http.Request) {

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

	wl, err := l.db.GetUserWorkouts(id)
	if err != nil {
		http.Error(rw, fmt.Sprintf("error returned: %s", err), http.StatusInternalServerError)
	}
	err = ToJSON(wl, rw) //Encode the list from structs to JSON objects
	if err != nil {      //if json can not be encoded return an error and log the error while also returning out of the function
		http.Error(rw, "Unable to encode JSON object", http.StatusInternalServerError)
		return
	}
}

func (l *Aer_Log) GetSingleWorkout(rw http.ResponseWriter, r *http.Request) {

	//swagger stuff here

	Athlete_ID := getAthleteID(r)
	Workout_ID, err := getWorkoutID(r)

	if err != nil {
		http.Error(rw, fmt.Sprintf("Error in converting string to int: %s", err), http.StatusInternalServerError)
		return
	}

	w, err := l.db.GetSingleWorkout(Athlete_ID, Workout_ID)

	if err != nil {
		http.Error(rw, fmt.Sprintf("Error in getting workout: %s", err), http.StatusBadRequest)
		return
	}

	err = ToJSON(w, rw)

	if err != nil {
		http.Error(rw, fmt.Sprintf("Error in serializing workout: %s", err), http.StatusBadRequest)
	}
}

func (l *Aer_Log) GetAthlete(rw http.ResponseWriter, r *http.Request) {

	//swagger here

	athlete_id := getAthleteID(r)
	athlete, err := l.db.GetAthlete(athlete_id)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Error in getting athlete: %s", err), http.StatusBadRequest)
		return
	}
	ToJSON(athlete, rw)
}

func (l *Aer_Log) GetAllAthletes(rw http.ResponseWriter, r *http.Request) {

	//swagger here

	athletes, err := l.db.GetAllAthletes()
	if err != nil {
		http.Error(rw, fmt.Sprintf("Error getting all athletes: %s", err), http.StatusBadRequest)
		return
	}
	fmt.Println("am i showing up")
	ToJSON(athletes, rw)

}
