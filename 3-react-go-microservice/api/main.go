package main

import (
	"log"
	"net/http"
)

var results []FizzBuzzResult

func main() {
	log.Print("Server listing on http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
