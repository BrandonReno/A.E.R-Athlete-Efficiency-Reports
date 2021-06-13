package controllers

import (
	"fmt"
	"net/http"
	"github.com/BrandonReno/A.E.R/models"
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
	err := l.db.CreateWorkout(&workout)
	if err != nil{
		l.l.Printf("Error: Could not create workout: %s", err)
		http.Error(rw, fmt.Sprintf("Error in creating workout: %s", err), http.StatusInternalServerError)
		return
	}

	athlete, err := l.db.GetAthlete(workout.Athlete_ID)

	if err != nil{
		l.l.Printf("Error: Could not find athlete: %s", err)
		http.Error(rw, fmt.Sprintf("Error could not find athlete: %s", err), http.StatusInternalServerError)
		return
	}

	err = l.db.UpdateEfficiency(&athlete)
	if err != nil{
		l.l.Printf("Error: Could not update efficiency score: %s", err)
		http.Error(rw, fmt.Sprintf("Error could not update efficiency score: %s", err), http.StatusInternalServerError)
		return

	}
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
	err := l.db.AddAthlete(&athlete)

	if err != nil{
		l.l.Printf("Error: Could not create athlete")
		http.Error(rw, fmt.Sprintf("Error in creating athlete: %s", err), http.StatusBadRequest)
		return
	}
}