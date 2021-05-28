package routes

import (
	"net/http"
	"github.com/BrandonReno/A.E.R/controllers"
	"github.com/gorilla/mux"
)

func initAthleteSR(r *mux.Router, l *controllers.Workout_Log){
	for _, route := range athleteRoutes{
		sr := r.NewRoute().Subrouter()
		sr.Methods(route.Request).Path(route.Pattern).Handler(route.Handler)
		if route.Request == http.MethodPost || route.Request == http.MethodPut{
			sr.Use(l.MiddlewarAthleteValidation)
		}
	}
}

var athleteRoutes = Routes{
	Route{
		Request: http.MethodGet,
		Pattern: "/athletes/{athlete_id:[[:alnum:]]+}",
		Handler: controllers.GetAthlete,
	},

	Route{
		Request: http.MethodPost,
		Pattern: "/athletes",
		Handler: controllers.CreateAthlete,
	},

	Route{
		Request: http.MethodPut,
		Pattern: "/athletes/{athlete_id:[[:alnum:]]+}",
		Handler: controllers.UpdateAthlete,
	},

	Route{
		Request: http.MethodDelete,
		Pattern: "/athletes/{athlete_id:[[:alnum:]]+}",
		Handler: controllers.DeleteAthlete,
	},
}
