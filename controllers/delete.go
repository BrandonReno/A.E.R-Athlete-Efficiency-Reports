package controllers

import (
	"fmt"
	"net/http"
)

// Delete a workout from the database
func (l * Aer_Log) DeleteWorkout(rw http.ResponseWriter, r *http.Request){
	// swagger:route DELETE /workouts/wid{id} workouts deleteWorkout
    //
    // Deletes a specified workout by id
    //
    //      Schemes: http
    //		
	//		Responses:
	//			201: noContent
	//			404: badRequest
	athlete_id := getAthleteID(r)
	workout_id, err := getWorkoutID(r)
	if err != nil{
		http.Error(rw, fmt.Sprintf("Error getting workout ID: %s", err), http.StatusBadRequest)
		return
	}

	err = l.db.DeleteWorkout(athlete_id, workout_id)
	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in deleting workout: %s", err), http.StatusBadRequest)
		return
	}
}


func (l * Aer_Log) DeleteAthlete(rw http.ResponseWriter, r *http.Request){


	//swagger

	athlete_id := getAthleteID(r)
	err := l.db.DeleteAthlete(athlete_id)
	if err != nil{
		http.Error(rw, fmt.Sprintf("Error deleting athlete: %s", err), http.StatusBadRequest)
		return
	}
	
}
