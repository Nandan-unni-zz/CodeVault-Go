package main

import (
	"net/http"

	"./routes"
	"./utils"
)

func main() {
	router := routes.Router()
	http.Handle("/", router)
	utils.LoadTemplates("templates/*.html")
	utils.LoadStaticFiles("./static")
	print("Server running at http://localhost:5000\nPress CTRL+C to quit")
	http.ListenAndServe(":5000", nil)
}
