package main

import (
	"net/http"
)

func main() {
	go println("Serving at http://localhost:8000/")
	panic(http.ListenAndServe(":8000", http.FileServer(http.Dir("."))))
}
