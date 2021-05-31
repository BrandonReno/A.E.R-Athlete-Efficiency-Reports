package controllers

import (
	"fmt"
	"net/http"
	"github.com/BrandonReno/A.E.R/models"
)

// Update a workout in the database
func (l *Aer_Log) UpdateWorkout(rw http.ResponseWriter, r *http.Request) {

	// swagger:route PUT /workouts/wid{id} workouts updateWorkout
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
	err = l.db.UpdateWorkout(&workout)

	if err != nil{
		http.Error(rw, fmt.Sprintf("Error updating workout: %s", err), http.StatusInternalServerError)
		return
	}
}

func (l *Aer_Log) UpdateAthlete(rw http.ResponseWriter, r *http.Request){


	//swagger
	athlete_id := getAthleteID(r)
	athlete := r.Context().Value(KeyCtx{}).(models.Athlete)
	athlete.Athlete_ID = athlete_id
	err := l.db.UpdateAthlete(&athlete)
	if err != nil{
		http.Error(rw, fmt.Sprintf("Error in updating athlete: %s", err), http.StatusBadRequest)
		return
	}

}