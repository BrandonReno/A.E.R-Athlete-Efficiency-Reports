package controllers

import (
	"fmt"
	"net/http"

	"github.com/BrandonReno/A.E.R/models"
	"github.com/BrandonReno/A.E.R/pooling"
)

//lists all workouts from all athletes
func (l *Aer_Log) GetAllWorkouts(rw http.ResponseWriter, r *http.Request){
	// swagger:route GET / workouts getAllWorkouts
	//
	// Lists workouts from all athletes accross the service
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//			200: workoutsResponse

	workout_channel := make(chan []models.Workout, 1)
	task := func() error{ return l.db.GetAllWorkouts(workout_channel)}

	job := pooling.Job{Name: "Get all Workouts", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)

	l.collector.EnqueJob(&job)

	wl := <- workout_channel

	err := ToJSON(wl, rw)

	if err != nil{
		l.l.Printf("Error: Could not serialize workouts: %s", err)
		http.Error(rw, fmt.Sprintf("Error serializing workouts: %s", err), http.StatusBadRequest)
		return
	}
}


// Gets all workouts in the system for an athlete
func (l *Aer_Log) GetWorkouts(rw http.ResponseWriter, r *http.Request) {

	// swagger:route GET /athletes/{athlete_id}/workouts workouts getUserWorkouts
	//
	// Lists workouts from a particular registered athlete
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//			200: workoutsResponse
	//			404: badRequest

	id := getAthleteID(r)

	wc := make(chan []models.Workout, 1)
	task := func() error { return l.db.GetUserWorkouts(id, wc)}

	job := pooling.Job{Name: "Get Athlete Workouts", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(&job)

	wl := <- wc
	

	err := ToJSON(wl, rw) //Encode the list from structs to JSON objects
	if err != nil {      //if json can not be encoded return an error and log the error while also returning out of the function
		l.l.Printf("Error: Could not encode to JSON: %s", err)
		http.Error(rw, "Unable to encode JSON object", http.StatusInternalServerError)
		return
	}
}

// Get a single workout in the system
func (l *Aer_Log) GetSingleWorkout(rw http.ResponseWriter, r *http.Request) {
	// swagger:route GET /athletes/{athlete_id}/workouts/{workout_id} workouts getSingleWorkout
	//
	// Gets a single workout from a specified athlete
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//			200: singleWorkout
	//			404: badRequest

	Athlete_ID := getAthleteID(r)
	Workout_ID, err := getWorkoutID(r)

	if err != nil {
		l.l.Printf("Error: Could not convert string to int: %s", err)
		http.Error(rw, fmt.Sprintf("Error in converting string to int: %s", err), http.StatusInternalServerError)
		return
	}

	wc := make(chan models.Workout, 1)
	task := func() error { return l.db.GetSingleWorkout(Athlete_ID, Workout_ID, wc)}

	job := pooling.Job{Name: "Get single Workout", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(&job)

	w := <- wc

	err = ToJSON(w, rw)

	if err != nil {
		l.l.Printf("Error: Could not encode to JSON: %s", err)
		http.Error(rw, fmt.Sprintf("Error in encoding workout to JSON: %s", err), http.StatusBadRequest)
		return
	}
}

// Get a single athlete in the system
func (l *Aer_Log) GetAthlete(rw http.ResponseWriter, r *http.Request) {
	// swagger:route GET /athletes/{athlete_id} athletes getAthlete
	//
	// Lists information from a specified athlete
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//			200: singleAthlete
	//			404: badRequest

	athlete_id := getAthleteID(r)

	ac := make(chan models.Athlete)
	task := func() error {return l.db.GetAthlete(athlete_id, ac)}

	job := pooling.Job{Name: "Get single Athlete", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(&job)

	athlete := <-ac

	err := ToJSON(athlete, rw)
	if err != nil {
		l.l.Printf("Error: Could not encode to JSON: %s", err)
		http.Error(rw, fmt.Sprintf("Error in encoding workout to JSON: %s", err), http.StatusBadRequest)
		return
	}
}

// Gets all athletes registered in the system
func (l *Aer_Log) GetAllAthletes(rw http.ResponseWriter, r *http.Request) {
	// swagger:route GET /athletes athletes listAthletes
	//
	// Lists all athletes registered in the system
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//			200: athletesResponse

	ac := make(chan []models.Athlete, 1)

	task := func() error{ return l.db.GetAllAthletes(ac)}

	job := pooling.Job{Name: "Get all Athletes", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(&job)

	athletes := <- ac
	
	err := ToJSON(athletes, rw)
	if err != nil {
		l.l.Printf("Error: Could not encode to JSON: %s", err)
		http.Error(rw, fmt.Sprintf("Error in encoding workout to JSON: %s", err), http.StatusBadRequest)
		return
	}
}

//Returns an athletes AER 
func (l *Aer_Log) GetAthleteEfficiency(rw http.ResponseWriter, r *http.Request){
	// swagger:route GET /athletes/{athlete_id}/aer efficiency getEfficiency
	//
	// Shows the athletes current AER statistics
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//			200: efficiencyResponse
	//			404: badRequest
	athleteID := getAthleteID(r)

	ac := make(chan models.Athlete, 1)

	task := func () error { return l.db.GetAthlete(athleteID, ac)}

	job := pooling.Job{Name: "Get Efficiency - Fetching Athlete", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(&job)

	athlete := <-ac


	ec := make(chan models.Efficiency)

	task = func() error { return l.db.GetEfficiency(&athlete, ec)}

	job = pooling.Job{Name: "Get Efficiency", Task: task}

	l.l.Printf("Job enqued to worker pool: %s", job.Name)
	l.collector.EnqueJob(&job)

	e := <- ec

	err := ToJSON(e, rw)

	if err != nil{
		l.l.Printf("Error: could not transfer efficiency to json: %s", err)
		http.Error(rw, fmt.Sprintf("Error inserializing efficiency: %s", err), http.StatusBadRequest)
		return
	}
}
