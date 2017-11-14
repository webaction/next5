package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"flag"
)

var (
	addr      = flag.String("addr", "127.0.0.1:9999", "http service address")
)

func main() {
	fmt.Println("Starting Server")

	router := mux.NewRouter()
	router.HandleFunc("/api/races", GetRaces).Methods("GET")

	httpErr := http.ListenAndServe(*addr, router)
	if httpErr != nil {
		log.Fatal("ListenAndServe: ", httpErr)
	}
}
