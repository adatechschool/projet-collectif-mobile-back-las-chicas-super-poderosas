package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "encoding/json"
	// "io/ioutil"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Salut la compagnie !")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/events", createEvents)
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

