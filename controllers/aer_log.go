package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"github.com/BrandonReno/A.E.R/pooling"
	"github.com/BrandonReno/A.E.R/services"
	"github.com/gorilla/mux"
)

type Aer_Log struct{
	l *log.Logger //Hold a logger that can write errors to the log output
	db *services.DB
	collector *pooling.Collector
}

type KeyCtx struct{} // used for storing context in middleware

func New(l *log.Logger, d *services.DB, collector *pooling.Collector) *Aer_Log{
	return &Aer_Log{l, d, collector} //creates and returns a new reference to the aer_Log
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

func ToJSON(i interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(i) // Create a new encoder and encode the current Workout_Feed to json. Returns an error just in case
}

func FromJSON(i interface{}, r io.Reader) error{
	return json.NewDecoder(r).Decode(i) // Create a new decoder and decode the request body to json. Returns an error just in case
}
