package routes

import (
	"net/http"
	"github.com/BrandonReno/A.E.R/handlers"
	"github.com/gorilla/mux"
)

func initWorkoutSR(r *mux.Router, l *handlers.Workout_Log){
	for _, route := range workoutRoutes{
		sr := r.NewRoute().Subrouter()
		sr.Methods(route.Request).Path(route.Pattern).Handler(route.Handler)
		if route.Request == http.MethodPost || route.Request == http.MethodPut{
			sr.Use(l.MiddlewareWorkoutValidation)
		}
	}
}

var workoutRoutes = Routes{
	Route{
		Request: http.MethodGet,
		Pattern: "/workouts/{athlete_id:[[:alnum:]]+}",
		Handler: handlers.GetWorkouts,
	},

	Route{
		Request: http.MethodGet,
		Pattern: "/workouts/{athlete_id:[[:alnum:]]+}/{workout_id:[0-9]+}",
		Handler: handlers.GetSingleWorkout,
	},

	Route{
		Request: http.MethodPost,
		Pattern: "/workouts/{athlete_id:[[:alnum:]]+}",
		Handler: handlers.AddWorkout,
	},

	Route{
		Request: http.MethodPut,
		Pattern: "/workouts/{athlete_id:[[:alnum:]]+}/{workout_id:[0-9]+}",
		Handler: handlers.UpdateWorkout,
	},

	Route{
		Request: http.MethodDelete,
		Pattern: "/workouts/{athlete_id:[[:alnum:]]+}/{workout_id:[0-9]+}",
		Handler: handlers.DeleteWorkout,
	},
}

	