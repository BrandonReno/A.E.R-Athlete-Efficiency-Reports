package handlers

import(
	"net/http"
	"github.com/BrandonReno/A.E.R/data"
	"fmt"
	"context"
)

//Middleware below, called before the subrouters handlerfunc. Example, when subrouter matches a POST verb middleware is called and then post
func (w *Workout_Log) MiddlewareWorkoutValidation(next http.Handler) http.Handler{ 
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){ 
		workout := data.Workout{} //Create a blank Workout 
		err := data.FromJSON(&workout, r.Body) //using the io reader of the request body read the json r.body and decode it to a workout.
		
		ID := getAthleteID(r)
		workout.Athlete_ID = ID
		
		if err != nil{ // if an error occurs while deserializing the workout from json print to the log and raise the http error
			w.l.Println("Error deserializing product")
			http.Error(rw, fmt.Sprintf("Unable to deserialize JSON object: %s", err), http.StatusInternalServerError)
			return
		}

		err = workout.Validate_Workout() //validate the workout for security
		if err != nil{ // if there is an error validating report which field is raising the error and print to log
			w.l.Println("Error validating product")
			http.Error(rw, fmt.Sprintf("Unable to validate JSON object: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyCtx{}, workout) //Store the new workout in the context map with key: KeyWorkout{} and value workout
		r = r.WithContext(ctx) //store the new context on the current http.Request context 

		next.ServeHTTP(rw, r) //Continue on to the next handle with the new context in the Request
	})
}

func (w *Workout_Log) MiddlewarAthleteValidation(next http.Handler) http.Handler{ 
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){ 
		athlete := data.Athlete{} //Create a blank athlete 
		err := data.FromJSON(&athlete, r.Body) //using the io reader of the request body read the json r.body and decode it to an athlete.

		fmt.Printf("%+v", athlete)
		
		if err != nil{ // if an error occurs while deserializing the athlete from json print to the log and raise the http error
			w.l.Println("Error deserializing product")
			http.Error(rw, fmt.Sprintf("Unable to deserialize JSON object: %s", err), http.StatusInternalServerError)
			return
		}

		err = athlete.Validate_Athlete() //validate the athlete for security
		if err != nil{ // if there is an error validating report which field is raising the error and print to log
			w.l.Println("Error validating product")
			w.l.Println(err)
			http.Error(rw, fmt.Sprintf("Unable to validate JSON object: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyCtx{}, athlete) //Store the new athlete in the context map with key: Keyathlete{} and value athlete
		r = r.WithContext(ctx) //store the new context on the current http.Request context 

		next.ServeHTTP(rw, r) //Continue on to the next handle with the new context in the Request
	})
}