package controllers

import (
	"fmt"
	"net/http"

	"github.com/BrandonReno/A.E.R/pooling"
)

// Delete a workout from the database
func (l * Aer_Log) DeleteWorkout(rw http.ResponseWriter, r *http.Request){
	// swagger:route DELETE /athletes/{athlete_id}/{workout_id} workouts deleteWorkout
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
	
	task := func() error { return l.db.DeleteWorkout(athlete_id, workout_id)}

	job := pooling.Job{Name: "Delete Workout", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(&job)
}

//Delete an athlete from the system
func (l * Aer_Log) DeleteAthlete(rw http.ResponseWriter, r *http.Request){
	// swagger:route DELETE /athletes/{athlete_id} athletes deleteAthlete
    //
    // Deletes a specified athlete by id
    //
    //      Schemes: http
    //		
	//		Responses:
	//			201: noContent
	//			404: badRequest

	athlete_id := getAthleteID(r)
	task := func() error{ return l.db.DeleteAthlete(athlete_id) }

	job := pooling.Job{Name: "Delete Athlete", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(&job)
}
