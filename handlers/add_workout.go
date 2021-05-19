package handlers

import(
	"net/http"
	"github.com/BrandonReno/Workouts/data"
)

// Add a workout to the database
func AddWorkout(rw http.ResponseWriter, r *http.Request){
	// swagger:route POST / workouts addWorkout
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

	workout := r.Context().Value(KeyWorkout{}).(data.Workout) //recieve the stored body from the request context and store the struct in a new workout object
	data.AddWorkout(&workout) //Add the workout to the list
}