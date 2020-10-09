package main

import (
	"fmt"
	"net/http"
)

func helloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Hello World from Go http<h1>")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}
