package routes

import (
	"../handlers"
	"github.com/gorilla/mux"
)

// Router : router
func Router() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", handlers.FeedGetHandler()).Methods("GET")
	route.HandleFunc("/", handlers.FeedPostHandler()).Methods("POST")
	route.HandleFunc("/vote/", handlers.VoteGetHandler()).Methods("GET")
	route.HandleFunc("/vote/", handlers.VotePostHandler()).Methods("POST")
	route.HandleFunc("/result/", handlers.ResultGetHandler()).Methods("GET")
	return route
}
