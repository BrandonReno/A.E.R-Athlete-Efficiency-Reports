package handlers

import (
	"net/http"
	"strconv"
	"github.com/BrandonReno/Workouts/data"
	"github.com/gorilla/mux"
)

// Delete a workout from the database
func DeleteWorkout(rw http.ResponseWriter, r *http.Request){
	// swagger:route DELETE /{id} workouts deleteWorkout
    //
    // Deletes a specified workout by id
    //
    //      Schemes: http
    //		
	//		Responses:
	//			201: noContent
	//			404: badRequest
    
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil{
		http.Error(rw, "Could not convert ID to string", http.StatusInternalServerError)
		return 
	}

	err = data.DeleteWorkout(id)

	if err != nil{
		http.Error(rw, "Could not delete workout", http.StatusNotFound)
		return
	}
}
