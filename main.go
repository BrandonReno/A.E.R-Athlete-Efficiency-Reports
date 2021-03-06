package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/BrandonReno/A.E.R/controllers"
	"github.com/BrandonReno/A.E.R/pooling"
	"github.com/BrandonReno/A.E.R/routes"
	"github.com/BrandonReno/A.E.R/services"
)

const Port = ":9090"
const Workers = 10

func main() {

	l := log.New(os.Stdout, "Workout-api", log.LstdFlags) //Create an instance of logger to use for Handlers

	db_user, db_pass, db_host, db_port, db_db := 
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB") 

	db := services.DB{}

	if err := db.OpenDBConnection(db_user, db_pass, db_host, db_db, db_port); err != nil {
		l.Fatal(err)
	}

	//Initialize the collector and begin listening for jobs
	collector := pooling.StartDispatcher(Workers, l)
	env := controllers.New(l, &db, collector) //Handler for Workout, gets a log to write errors and such

	muxRouter := routes.NewRouter(env)

	//Create the server :"s"
	s := http.Server{
		Addr:         Port,              //Port 9090
		Handler:      muxRouter,         //Handler is our created mux server_mux
		ReadTimeout:  10 * time.Second,   //Read timeout is 1 second
		WriteTimeout: 10 * time.Second,   //Write timeout is 1 second
		IdleTimeout:  120 * time.Second, //Idle timeout is 120 seconds
	}

	l.Println("Server is up and running")
	// Listen and serve in a go routine to allow for graceful shutdown
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatalln(err) // if listen and serve returns an error log the error
		}
	}()

	//Graceful shutdown set up below

	sigChan := make(chan os.Signal)      //make a channel which recieves os.Signal
	signal.Notify(sigChan, os.Interrupt) //Notify the channel if there is an interrupt
	signal.Notify(sigChan, os.Kill)      //Notify the channel if there is a kill call

	sig_result := <-sigChan //Send channel Signal output to result to log the reasoning for shutdown
	l.Println("Shutdown initiated with ", sig_result)
	collector.EndProcesses() // Stop all workers and stop accepting jobs
	l.Println("All workers successfully have been stopped")
	db.CloseDBConnection() //close database connection
	l.Println("Postgres connecttion successfully closed")

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second) //Give system 30 seconds to complete handlers
	defer cancel()
	s.Shutdown(tc) //Shutdown server gracefully

}
