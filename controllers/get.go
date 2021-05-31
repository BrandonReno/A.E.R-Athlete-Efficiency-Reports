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
		l.l.Printf("Error: Could not obtain user workouts: %s", err)
		http.Error(rw, fmt.Sprintf("Error getting user workouts: %s", err), http.StatusBadRequest)
		return
	}
	err = ToJSON(wl, rw) //Encode the list from structs to JSON objects
	if err != nil {      //if json can not be encoded return an error and log the error while also returning out of the function
		l.l.Printf("Error: Could not encode to JSON: %s", err)
		http.Error(rw, "Unable to encode JSON object", http.StatusInternalServerError)
		return
	}
}

// Get a single workout in the system
func (l *Aer_Log) GetSingleWorkout(rw http.ResponseWriter, r *http.Request) {
	// swagger:route GET /workouts/athleteID{aid}/workoutID{ID} workouts getSingleWorkout
	//
	// Gets a single workout from a specified athlete
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//			200: singleWorkout

	Athlete_ID := getAthleteID(r)
	Workout_ID, err := getWorkoutID(r)

	if err != nil {
		l.l.Printf("Error: Could not convert string to int: %s", err)
		http.Error(rw, fmt.Sprintf("Error in converting string to int: %s", err), http.StatusInternalServerError)
		return
	}

	w, err := l.db.GetSingleWorkout(Athlete_ID, Workout_ID)

	if err != nil {
		l.l.Printf("Error: Could not get workoutID, athleteID match: %s", err)
		http.Error(rw, fmt.Sprintf("Error in getting workout: %s", err), http.StatusBadRequest)
		return
	}

	err = ToJSON(w, rw)

	if err != nil {
		l.l.Printf("Error: Could not encode to JSON: %s", err)
		http.Error(rw, fmt.Sprintf("Error in encoding workout to JSON: %s", err), http.StatusBadRequest)
		return
	}
}

// Get a single athlete in the system
func (l *Aer_Log) GetAthlete(rw http.ResponseWriter, r *http.Request) {
	// swagger:route GET /athletes/athleteID{aid} athletes getAthlete
	//
	// Lists information from a specified athlete
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//			200: singleAthlete

	athlete_id := getAthleteID(r)
	athlete, err := l.db.GetAthlete(athlete_id)
	if err != nil {
		l.l.Printf("Error: could not get athlete from athlete_id %s : %s", athlete_id, err)
		http.Error(rw, fmt.Sprintf("Error in getting athlete: %s", err), http.StatusBadRequest)
		return
	}
	err = ToJSON(athlete, rw)
	if err != nil {
		l.l.Printf("Error: Could not encode to JSON: %s", err)
		http.Error(rw, fmt.Sprintf("Error in encoding workout to JSON: %s", err), http.StatusBadRequest)
		return
	}
}

// Gets all athletes registered in the system
func (l *Aer_Log) GetAllAthletes(rw http.ResponseWriter, r *http.Request) {
	// swagger:route GET / athletes listathletes
	//
	// Lists all athletes registered in the system
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//			200: athletesResponse

	athletes, err := l.db.GetAllAthletes()
	if err != nil {
		l.l.Printf("Error: Could not retrieve list of athletes: %s", err)
		http.Error(rw, fmt.Sprintf("Error getting all athletes: %s", err), http.StatusBadRequest)
		return
	}
	
	err = ToJSON(athletes, rw)
	if err != nil {
		l.l.Printf("Error: Could not encode to JSON: %s", err)
		http.Error(rw, fmt.Sprintf("Error in encoding workout to JSON: %s", err), http.StatusBadRequest)
		return
	}
}
