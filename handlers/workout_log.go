package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Workout_Log struct{
	l *log.Logger //Hold a logger that can write errors to the log output
}

type KeyCtx struct{} // used for storing context in middleware

func New(l *log.Logger) *Workout_Log{
	return &Workout_Log{l} //creates and returns a new reference to the Workout_Log
}

func getAthleteID(r *http.Request) string{
	params := mux.Vars(r)
	ID := params["athlete_id"]
	return ID
}

func getWorkoutID(r *http.Request) (int, error){
	params := mux.Vars(r)
	ID, err := strconv.Atoi(params["workout_id"])
	return ID, err
}
