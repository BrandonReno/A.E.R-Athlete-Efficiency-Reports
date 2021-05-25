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

type KeyWorkout struct{} // used for storing context in middleware

func New(l *log.Logger) *Workout_Log{
	return &Workout_Log{l} //creates and returns a new reference to the Workout_Log
}

func getProductID(r *http.Request) int{
	params := mux.Vars(r)
	ID, err := strconv.Atoi(params["id"])

	if err != nil{
		panic(err) // will never happen but just for defense
	}

	return ID

}
