package main

import (
	"net/http"

	"./routes"
	"./utils"
)

func main() {
	print("Connected to DataBase.\n")
	utils.LoadTemplates("templates/*.html")
	print("Loaded templates.\n")
	router := routes.Router()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	print("Loaded static files.\n")
	http.Handle("/", router)
	print("Server running at http://localhost:5000\n")
	http.ListenAndServe(":5000", nil)
}
