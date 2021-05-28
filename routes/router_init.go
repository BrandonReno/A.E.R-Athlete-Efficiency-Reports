package routes

import (
	"net/http"
	"github.com/BrandonReno/A.E.R/handlers"
	"github.com/gorilla/mux"
)

//Create type route to then create subroutes for athlete and workout
type Route struct {
	Request string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

func NewRouter(l *handlers.Workout_Log) *mux.Router{
	serve_mux := mux.NewRouter()
	initAthleteSR(serve_mux, l)
	initWorkoutSR(serve_mux, l)
	return serve_mux
}