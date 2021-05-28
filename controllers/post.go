package controllers

import (
	"fmt"
	"net/http"
	"github.com/BrandonReno/A.E.R/models"
	"github.com/BrandonReno/A.E.R/services"
)

// Add a workout to the database
func AddWorkout(rw http.ResponseWriter, r *http.Request){
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
	err := services.CreateWorkout(&workout)
	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in creating workout: %s", err), http.StatusInternalServerError)
		return
	}
	

}

func CreateAthlete(rw http.ResponseWriter, r *http.Request){

	//swagger

	athlete := r.Context().Value(KeyCtx{}).(models.Athlete)
	err := services.AddAthlete(&athlete)

	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in creating athlete: %s", err), http.StatusBadRequest)
		return
	}
}