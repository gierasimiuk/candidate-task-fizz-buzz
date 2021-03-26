package main

import (
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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fizzbuzz", fizzbuzz).Methods("POST")
	log.Print("Server listing on http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
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

	fizzBuzzResults, _ := fizzBuzz(start, end)
	log.Print(fizzBuzzResults)

	// TODO Send result back in JSON format
	for _, s := range fizzBuzzResults {
		input := strconv.Itoa(s.Input)
		fmt.Fprintf(w, input+" "+s.Output+"<br>")
	}
}
