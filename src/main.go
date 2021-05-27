package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/BrandonReno/A.E.R/data"
	"github.com/BrandonReno/A.E.R/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

const Port = ":9090"


func main(){

	l := log.New(os.Stdout, "Workout-api", log.LstdFlags) //Create an instance of logger to use for Handlers
	
	wl := handlers.New(l) //Handler for Workout, gets a log to write errors and such

	server_mux := mux.NewRouter() //Create a new mux router to handle RESTful services
	
	WGetSrouter :=server_mux.Methods(http.MethodGet).Subrouter() //Create a subrouter of router server_mux just for get requests
	WGetSrouter.HandleFunc("/workouts/{athlete_id:[[:alnum:]]+}", handlers.GetWorkouts)
	WGetSrouter.HandleFunc("/workouts/{athlete_id:[[:alnum:]]+}/{workout_id:[0-9]+}", handlers.GetSingleWorkout)
	

	WDeleteSrouter := server_mux.Methods(http.MethodDelete).Subrouter() // Create a subrouter of router server_mux for delete requests
	WDeleteSrouter.HandleFunc("/workouts/{athlete_id:[[:alnum:]]+}/{workout_id:[0-9]+}", handlers.DeleteWorkout)

	WPutSrouter := server_mux.Methods(http.MethodPut).Subrouter() //Create a subrouter of router server_mux just for put requests
	WPutSrouter.HandleFunc("/workouts/{athlete_id:[[:alnum:]]+}/{workout_id:[0-9]+}", handlers.UpdateWorkout)
	WPutSrouter.Use(wl.MiddlewareWorkoutValidation) //Add middleware, step before Handlefunc

	WPostSrouter := server_mux.Methods(http.MethodPost).Subrouter()  //Create a subrouter of router server_mux just for post requests
	WPostSrouter.HandleFunc("/workouts/{athlete_id:[[:alnum:]]+}", handlers.AddWorkout)
	WPostSrouter.Use(wl.MiddlewareWorkoutValidation) //Add middleware, step before Handlefunc
	
	

	AGetSrouter := server_mux.Methods(http.MethodGet).Subrouter()
	AGetSrouter.HandleFunc("/athlete/{athlete_id:[[:alnum:]]+}", handlers.GetAthlete)
	
	APostSrouter := server_mux.Methods(http.MethodPost).Subrouter()
	APostSrouter.HandleFunc("/athlete", handlers.CreateAthlete)
	APostSrouter.Use(wl.MiddlewarAthleteValidation)

	APutSrouter := server_mux.Methods(http.MethodPut).Subrouter()
	APutSrouter.HandleFunc("/athlete/{athlete_id:[[:alnum:]]+}", handlers.UpdateAthlete)
	APutSrouter.Use(wl.MiddlewarAthleteValidation)

	ADeleteSrouter := server_mux.Methods(http.MethodDelete).Subrouter()
	ADeleteSrouter.HandleFunc("/athlete/{athlete_id:[[:alnum:]]+}", handlers.DeleteAthlete)



	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	WGetSrouter.Handle("/docs", sh) //Set up the GetSrouter to also handle the docs
	WGetSrouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./"))) //Serve the swagger.yaml file on the server

	//Create the server :"s"
	s := http.Server{
		Addr: Port, //Port 9090
		Handler: server_mux, //Handler is our created mux server_mux
		ReadTimeout: 1 * time.Second, //Read timeout is 1 second
		WriteTimeout: 1 * time.Second, //Write timeout is 1 second
		IdleTimeout: 120 * time.Second, //Idle timeout is 120 seconds
	}

	err := data.OpenDBConnection()

	if err != nil{
		l.Fatal(err)
	}

	// Listen and serve in a go routine to allow for graceful shutdown
	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Fatalln(err) // if listen and serve returns an error log the error
		}
	}()

	//Graceful shutdown set up below

	sigChan := make(chan os.Signal) //make a channel which recieves os.Signal
	signal.Notify(sigChan, os.Interrupt) //Notify the channel if there is an interrupt
	signal.Notify(sigChan, os.Kill) //Notify the channel if there is a kill call

	sig_result := <-sigChan //Send channel Signal output to result to log the reasoning for shutdown
	l.Println("Shutdown initiated with ", sig_result)
	data.DBConn.Close() //close database connection

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second) //Give system 30 seconds to complete handlers
	s.Shutdown(tc) //Shutdown server gracefully

}

