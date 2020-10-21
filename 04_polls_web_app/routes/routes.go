package routes

import (
	"../handlers"
	"github.com/gorilla/mux"
)

// Router : router
func Router() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", handlers.FeedHandler).Methods("GET", "POST")
	route.HandleFunc("/vote", handlers.VoteHandler).Methods("GET", "POST")
	route.HandleFunc("/result", handlers.ResultHandler).Methods("GET")
	return route
}
