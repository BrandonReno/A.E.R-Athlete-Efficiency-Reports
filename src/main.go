package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/BrandonReno/A.E.R/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

const Port = ":9090"


func main(){

	l := log.New(os.Stdout, "Workout-api", log.LstdFlags) //Create an instance of logger to use for Handlers
	
	wl := handlers.New(l) //Handler for Workout, gets a log to write errors and such

	server_mux := mux.NewRouter() //Create a new mux router to handle RESTful services
	
	GetSrouter :=server_mux.Methods(http.MethodGet).Subrouter() //Create a subrouter of router server_mux just for get requests
	GetSrouter.HandleFunc("/", handlers.GetWorkouts)

	DeleteSrouter := server_mux.Methods(http.MethodDelete).Subrouter() // Create a subrouter of router server_mux for delete requests
	DeleteSrouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteWorkout)

	PutSrouter := server_mux.Methods(http.MethodPut).Subrouter() //Create a subrouter of router server_mux just for put requests
	PutSrouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateWorkout)
	PutSrouter.Use(wl.MiddlewareWorkoutValidation) //Add middleware, step before Handlefunc

	PostSrouter := server_mux.Methods(http.MethodPost).Subrouter()  //Create a subrouter of router server_mux just for post requests
	PostSrouter.HandleFunc("/", handlers.AddWorkout)
	PostSrouter.Use(wl.MiddlewareWorkoutValidation) //Add middleware, step before Handlefunc

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	GetSrouter.Handle("/docs", sh) //Set up the GetSrouter to also handle the docs
	GetSrouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./"))) //Serve the swagger.yaml file on the server

	//Create the server :"s"
	s := http.Server{
		Addr: Port, //Port 9090
		Handler: server_mux, //Handler is our created mux server_mux
		ReadTimeout: 1 * time.Second, //Read timeout is 1 second
		WriteTimeout: 1 * time.Second, //Write timeout is 1 second
		IdleTimeout: 120 * time.Second, //Idle timeout is 120 seconds
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

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second) //Give system 30 seconds to complete handlers
	s.Shutdown(tc) //Shutdown server gracefully

}

