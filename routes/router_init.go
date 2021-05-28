package routes

import (
	"net/http"
	"github.com/BrandonReno/A.E.R/handlers"
	"github.com/go-openapi/runtime/middleware"
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

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	docs := serve_mux.Methods(http.MethodGet).Subrouter()
	docs.Handle("/docs", sh) //Set up the GetSrouter to also handle the docs
	docs.Handle("/swagger.yaml", http.FileServer(http.Dir("./"))) //Serve the swagger.yaml file on the server

	return serve_mux
}