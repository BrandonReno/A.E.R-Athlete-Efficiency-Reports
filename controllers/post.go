package controllers

import (
	"fmt"
	"net/http"
	"github.com/BrandonReno/A.E.R/models"
)

// Add a workout to the database
func (l *Aer_Log) AddWorkout(rw http.ResponseWriter, r *http.Request){
	// swagger:route POST /workouts workouts addWorkout
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
		http.Error(rw, fmt.Sprintf("Error in creating workout: %s", err), http.StatusInternalServerError)
		return
	}
	

}

func (l *Aer_Log) CreateAthlete(rw http.ResponseWriter, r *http.Request){

	//swagger

	athlete := r.Context().Value(KeyCtx{}).(models.Athlete)
	err := l.db.AddAthlete(&athlete)

	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in creating athlete: %s", err), http.StatusBadRequest)
		return
	}
}