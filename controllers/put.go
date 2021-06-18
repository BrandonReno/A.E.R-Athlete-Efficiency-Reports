package controllers

import (
	"fmt"
	"net/http"

	"github.com/BrandonReno/A.E.R/models"
	"github.com/BrandonReno/A.E.R/pooling"
)

// Update a workout in the database
func (l *Aer_Log) UpdateWorkout(rw http.ResponseWriter, r *http.Request) {
	// swagger:route PUT /athletes/{athlete_id}/workouts/{workout_id} workouts updateWorkout
    //
    // Updates a workout in the system based on the given ID
    //
	//      Produces:
	//      - application/json
	//
	//		Consumes:
	//		- application/json
    //
    //     	Schemes: http
	//
	// 		Responses:
	//			201 : noContent
	//			400 : verror
	//			404 : badRequest		
    workout := r.Context().Value(KeyCtx{}).(models.Workout)
	wid, err := getWorkoutID(r)

	if err != nil{
		http.Error(rw, fmt.Sprintf("Error getting workout id: %s", err), http.StatusInternalServerError)
		return
	}

	workout.Workout_ID = wid

	task := func() error { return l.db.UpdateWorkout(&workout)}

	job := pooling.Job{Name: "Update Workout", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)

	l.collector.EnqueJob(&job)
}

// Update an athlete in the database
func (l *Aer_Log) UpdateAthlete(rw http.ResponseWriter, r *http.Request){
	// swagger:route PUT /athletes/{athlete_id} athletes updateAthlete
    //
    // Updates an athlete in the system based on the given ID
    //
	//      Produces:
	//      - application/json
	//
	//		Consumes:
	//		- application/json
    //
    //     	Schemes: http
	//
	// 		Responses:
	//			201 : noContent
	//			400 : verror
	//			404 : badRequest

	athlete_id := getAthleteID(r)
	athlete := r.Context().Value(KeyCtx{}).(models.Athlete)
	athlete.Athlete_ID = athlete_id
	

	task := func() error{ return l.db.UpdateAthlete(&athlete)}

	job := pooling.Job{Name: "Update Athlete", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	
	l.collector.EnqueJob(&job)
}