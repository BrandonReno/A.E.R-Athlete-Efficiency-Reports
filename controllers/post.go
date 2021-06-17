package controllers

import (
	"net/http"
	"github.com/BrandonReno/A.E.R/models"
	"github.com/BrandonReno/A.E.R/pooling"
)

// Add a workout to the database
func (l *Aer_Log) AddWorkout(rw http.ResponseWriter, r *http.Request){
	// swagger:route POST /athletes/{athlete_id} workouts addWorkout
    //
    // Adds a new workout to the database
    //
    //     	Consumes:
    //     	- application/json
	//
    //     	Schemes: http
	//
	// 		Responses:
	//			201 : noContent
	//			400 : verror

	workout := r.Context().Value(KeyCtx{}).(models.Workout)

	task := func() error{ return l.db.CreateWorkout(&workout)}

	job := pooling.Job{Name: "Create Workout", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(&job)
}

// Creates a new athlete
func (l *Aer_Log) CreateAthlete(rw http.ResponseWriter, r *http.Request){
	// swagger:route POST /athletes athletes addAthlete
    //
    // Adds a new workout to the database
    //
    //     	Consumes:
    //     	- application/json
	//
    //     	Schemes: http
	//
	// 		Responses:
	//			201 : noContent
	//			400 : verror

	athlete := r.Context().Value(KeyCtx{}).(models.Athlete)

	task := func() error{return l.db.AddAthlete(&athlete)}

	
	job := &pooling.Job{Name: "Create Athlete", Task: task}
	
	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(job)
}