package routes

import (
	"net/http"
	"github.com/BrandonReno/A.E.R/controllers"
	"github.com/gorilla/mux"
)

func initWorkoutSR(r *mux.Router, l *controllers.Aer_Log){
	workoutRoutes := initWorkoutSlice(l)
	for _, route := range workoutRoutes{
		sr := r.NewRoute().Subrouter()
		sr.Methods(route.Request).Path(route.Pattern).Handler(route.Handler)
		if route.Request == http.MethodPost || route.Request == http.MethodPut{
			sr.Use(l.MiddlewareWorkoutValidation)
		}
	}
}
func initWorkoutSlice(l *controllers.Aer_Log) Routes{
	workoutRoutes := Routes{
		Route{
			Request: http.MethodGet,
			Pattern: "/workouts/{athlete_id:[[:alnum:]]+}",
			Handler: l.GetWorkouts,
		},

		Route{
			Request: http.MethodGet,
			Pattern: "/workouts/{athlete_id:[[:alnum:]]+}/{workout_id:[0-9]+}",
			Handler: l.GetSingleWorkout,
		},

		Route{
			Request: http.MethodPost,
			Pattern: "/workouts/{athlete_id:[[:alnum:]]+}",
			Handler: l.AddWorkout,
		},

		Route{
			Request: http.MethodPut,
			Pattern: "/workouts/{athlete_id:[[:alnum:]]+}/{workout_id:[0-9]+}",
			Handler: l.UpdateWorkout,
		},

		Route{
			Request: http.MethodDelete,
			Pattern: "/workouts/{athlete_id:[[:alnum:]]+}/{workout_id:[0-9]+}",
			Handler: l.DeleteWorkout,
		},
	}
	return workoutRoutes
}

	