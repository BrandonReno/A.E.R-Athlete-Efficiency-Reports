package routes

import (
	"net/http"
	"github.com/BrandonReno/A.E.R/controllers"
	"github.com/gorilla/mux"
)

func initAthleteSR(r *mux.Router, l *controllers.Aer_Log){
	athleteRoutes := initAthleteSlice(l)
	for _, route := range athleteRoutes{
		sr := r.NewRoute().Subrouter()
		sr.Methods(route.Request).Path(route.Pattern).Handler(route.Handler)
		if route.Request == http.MethodPost || route.Request == http.MethodPut{
			sr.Use(l.MiddlewarAthleteValidation)
		}
	}
}

func initAthleteSlice(l *controllers.Aer_Log) Routes{
	athleteRoutes := Routes{
		Route{
			Request: http.MethodGet,
			Pattern: "/",
			Handler: l.WelcomeRoute,
		},

		Route{
			Request: http.MethodGet,
			Pattern: "/athletes/{athlete_id:[[:alnum:]]+}",
			Handler: l.GetAthlete,
		},

		Route{
			Request: http.MethodPost,
			Pattern: "/athletes",
			Handler: l.CreateAthlete,
		},

		Route{
			Request: http.MethodPut,
			Pattern: "/athletes/{athlete_id:[[:alnum:]]+}",
			Handler: l.UpdateAthlete,
		},

		Route{
			Request: http.MethodDelete,
			Pattern: "/athletes/{athlete_id:[[:alnum:]]+}",
			Handler: l.DeleteAthlete,
		},

		Route{
			Request: http.MethodGet,
			Pattern: "/athletes",
			Handler: l.GetAllAthletes,
		},

		Route{
			Request: http.MethodGet,
			Pattern: "/athletes/{athlete_id:[[:alnum:]]+}/aer",
			Handler: l.GetAthleteEfficiency,
		},
	}
	return athleteRoutes
}
