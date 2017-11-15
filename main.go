package main

import (
	"fmt"
	"net/http"
	"log"
	"flag"
	"encoding/json"
	"time"
	"math/rand"
	"sort"

	"github.com/gorilla/mux"
	"github.com/icrowley/fake"
)

var (
	addr      = flag.String("addr", "127.0.0.1:9999", "http service address")
	races     []Race
	raceTypes []string
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

type bySuspend []Race

func (a bySuspend) Len() int      { return len(a) }
func (a bySuspend) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a bySuspend) Less(i, j int) bool {
	if a[i].Suspend < a[j].Suspend {
		return true
	}
	if a[i].Suspend > a[j].Suspend {
		return false
	}
	return a[i].Suspend < a[j].Suspend
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomTimestamp(i int) int64 {

	randomTime := time.Now().Unix() + (int64(i) * 600)

	randomNow := time.Unix(randomTime, 0).Unix()

	return randomNow
}

func buildDataSet() []Race {
	raceTypes = []string{"Thoroughbred", "Greyhounds", "Harness"}

	var race []Race
	for i := 0; i < 1000; i++ {
		idx := random(0, 2)
		suspend := randomTimestamp(i)
		name := fake.FullName()

		var competitors []Competitors
		for h := 1; h < random(4, 23); h++ {
			competitors = append(competitors, Competitors{Position: h, Name: fake.FullName()})
		}

		race = append(race, Race{Description: name,
			Suspend: suspend,
			EventID: random(10000, 50000),
			Details: Details{
				RaceType: raceTypes[idx],
				RaceNum:  i,
				Country:  "AUS",
				Meeting: Meeting{
					ID:   random(100, 9999),
					Name: name,
					Date: time.Unix(suspend, 0).String(),
				},
			},
			Competitors: competitors,
		})
	}

	sort.Sort(bySuspend(race))

	return race
}

func FilterRaces(vs []Race) []Race {
	vsf := make([]Race, 0)
	i := 0
	for _, v := range vs {
		if v.Suspend >= time.Now().Unix() && i < 5 {
			vsf = append(vsf, v)
			i++
		}
	}
	return vsf
}

func GetRaces(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/javascript")
	json.NewEncoder(writer).Encode(FilterRaces(races))
}

func main() {
	fmt.Println("Starting Server")

	races = buildDataSet()

	router := mux.NewRouter()
	router.HandleFunc("/api/races", GetRaces).Methods("GET")

	httpErr := http.ListenAndServe(*addr, router)
	if httpErr != nil {
		log.Fatal("ListenAndServe: ", httpErr)
	}
}
