package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

//Create type route to then create subroutes for athlete and workout
type Route struct {
	Request string
	Pattern string
	Handler http.HandlerFunc
}

func NewRouter() mux.Router{
	serve_mux := mux.NewRouter()


}