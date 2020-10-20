package main

import (
	"net/http"

	"./routes"
)

func main() {
	models.init()
	router = routes.Router()
	http.HandleFunc("", router)
	http.ListenAndServe(":5000", nil)
}
