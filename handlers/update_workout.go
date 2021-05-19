package handlers

import(
	"net/http"
	"github.com/BrandonReno/A.E.R/data"
	"github.com/gorilla/mux"
	"strconv"
)


// Update a workout in the database
func UpdateWorkout(rw http.ResponseWriter, r *http.Request) {

	// swagger:route PUT /{id} workouts updateWorkout
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
    
	params := mux.Vars(r) // get the list of parameters from the URI
	id,err := strconv.Atoi(params["id"]) // store the value of the 'id' URI value from the params map and convert it to an int
	if err != nil{  // if for some reason the string id can not convert to an int write an error the the response writer
		http.Error(rw, "Could not convert string ID to int", http.StatusBadRequest)
		return
	}
	
	workout := r.Context().Value(KeyWorkout{}).(data.Workout) // recieve the stored body from the request context and store the struct in a new workout object
	err = data.UpdateWorkout(id, &workout) // call the update data function which will return an error if the workout to be updated does not exist

	if err == data.ErrorWorkoutNotFound{ // If the error raised from Updateworkout is ErrorWorkout... then write the error to response writer
		http.Error(rw, "Workout not found", http.StatusNotFound)
		return
	}
	if err != nil{ // If the error raised from Updateworkout is not ErrorWorkout... then write the error to response writer
		http.Error(rw, "Error Updating Data", http.StatusNotFound)
		return
	}
}