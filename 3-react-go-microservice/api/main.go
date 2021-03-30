package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// Common library Used for routing with GoLang
	// Use the following command to install globally:
	// go get -u github.com/gorilla/mux
	"github.com/gorilla/mux"
)

func main() {
	router := routes()
	log.Print("Server listing on http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Build our router and set routes for the microservice
// then return the mux.Router instance
func routes() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fizzbuzz", fizzbuzz).Methods("POST")
	return router
}

func fizzbuzz(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	start, startErr := strconv.Atoi(r.FormValue("start"))
	end, endErr := strconv.Atoi(r.FormValue("end"))

	if startErr != nil {
		// TODO HANDLE ERROR
		log.Print(startErr)
		fmt.Fprintf(w, "Please enter numbers only")
		return
	}
	if endErr != nil {
		// TODO HANDLE ERROR
		log.Print(endErr)
		fmt.Fprintf(w, "Please enter numbers only")
		return
	}
	if start >= end {
		log.Print(endErr)
		fmt.Fprintf(w, "Ensure end number is greater than start")
		return
	}

	// Perform fizzbuzz calculation
	fizzBuzzResults, _ := fizzBuzz(start, end)

	// Send data back as JSON - ensuring header is set appropriately
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fizzBuzzResults)
}
