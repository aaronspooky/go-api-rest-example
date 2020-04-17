package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome home")
}

func main() {
	fmt.Print("Hello world!")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type event struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID: "1",
		Title: "Introduction to Golang",
		Description: "Come join us",
	},
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with event")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}