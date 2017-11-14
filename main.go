package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"flag"
	"encoding/json"
)

var (
	addr  = flag.String("addr", "127.0.0.1:9999", "http service address")
	races []Race
)

type Race struct {
	Description string        `json:"description"`
	Suspend     int64         `json:"suspend"`
	EventID     int           `json:"event_id"`
	Details     Details       `json:"details"`
	Competitors []Competitors `json:"competitors"`
}

type Details struct {
	RaceType string  `json:"race_type"`
	Country  string  `json:"country"`
	RaceNum  int     `json:"race_num"`
	Meeting  Meeting `json:"meeting"`
}

type Meeting struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

type Competitors struct {
	Position int    `json:"position"`
	Name     string `json:"name"`
}

func GetRaces(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/javascript")
	json.NewEncoder(writer).Encode(races)
}

func main() {
	fmt.Println("Starting Server")

	router := mux.NewRouter()
	router.HandleFunc("/api/races", GetRaces).Methods("GET")

	httpErr := http.ListenAndServe(*addr, router)
	if httpErr != nil {
		log.Fatal("ListenAndServe: ", httpErr)
	}
}
