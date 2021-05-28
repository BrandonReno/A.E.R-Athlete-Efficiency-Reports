package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/BrandonReno/A.E.R/services"
	"github.com/BrandonReno/A.E.R/controllers"
	"github.com/BrandonReno/A.E.R/routes"
)

const Port = ":9090"


func main(){

	l := log.New(os.Stdout, "Workout-api", log.LstdFlags) //Create an instance of logger to use for Handlers
	
	wl := controllers.New(l) //Handler for Workout, gets a log to write errors and such

	muxRouter := routes.NewRouter(wl)

	//Create the server :"s"
	s := http.Server{
		Addr: Port, //Port 9090
		Handler: muxRouter, //Handler is our created mux server_mux
		ReadTimeout: 1 * time.Second, //Read timeout is 1 second
		WriteTimeout: 1 * time.Second, //Write timeout is 1 second
		IdleTimeout: 120 * time.Second, //Idle timeout is 120 seconds
	}

	err := services.OpenDBConnection()

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
	services.DBConn.Close() //close database connection

	tc, cancel := context.WithTimeout(context.Background(), 30 * time.Second) //Give system 30 seconds to complete handlers
	defer cancel()
	s.Shutdown(tc) //Shutdown server gracefully

}

