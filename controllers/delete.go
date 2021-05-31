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
		l.l.Printf("Error: Unable to convert workout id to string: %s", err)
		http.Error(rw, fmt.Sprintf("Error getting workout ID: %s", err), http.StatusBadRequest)
		return
	}

	err = l.db.DeleteWorkout(athlete_id, workout_id)
	if err != nil{
		l.l.Printf("Error: Unable to delete athlete, can not find either workout id or athlete id: %s", err)
		http.Error(rw, fmt.Sprintf("Error in deleting workout: %s", err), http.StatusBadRequest)
		return
	}
}

//Delete an athlete from the system
func (l * Aer_Log) DeleteAthlete(rw http.ResponseWriter, r *http.Request){
	// swagger:route DELETE /athletes/athlete{id} athletes deleteAthlete
    //
    // Deletes a specified athlete by id
    //
    //      Schemes: http
    //		
	//		Responses:
	//			201: noContent
	//			404: badRequest

	athlete_id := getAthleteID(r)
	err := l.db.DeleteAthlete(athlete_id)
	if err != nil{
		l.l.Printf("Error: Could not delete athlete, athlete id may be incorrect: %s", err)
		http.Error(rw, fmt.Sprintf("Error deleting athlete: %s", err), http.StatusBadRequest)
		return
	}
}
