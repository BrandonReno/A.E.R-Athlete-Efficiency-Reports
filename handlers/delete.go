package handlers

import (
	"net/http"
)

// Delete a workout from the database
func DeleteWorkout(rw http.ResponseWriter, r *http.Request){
	// swagger:route DELETE /workouts/wid{id} workouts deleteWorkout
    //
    // Deletes a specified workout by id
    //
    //      Schemes: http
    //		
	//		Responses:
	//			201: noContent
	//			404: badRequest
    
}
