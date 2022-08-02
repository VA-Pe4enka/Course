package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Station struct {
	ID          string `json:"id"`
	OpeningTime string `json:"opening_time"`
	ClosinTime  string `json:"closin_time"`
}

type Schedule struct {
	StationID     string `json:"station_id"`
	Num           string `json:"num"`
	Direction     string `json:"direction"`
	ArrivalTime   string `json:"arrival_time"`
	DepartureTime string `json:"departure_time"`
}

func PostStation(w http.ResponseWriter, r *http.Request) {

	postStation()

	fmt.Println("Station has been successfully created and posted")

}

func GetStation(w http.ResponseWriter, r *http.Request) {

	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	stationData, scheduleForStation := getStation(id)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(stationData)
	json.NewEncoder(w).Encode(scheduleForStation)

}
func PostSchedule(w http.ResponseWriter, r *http.Request) {

	postSchedule()

	fmt.Println("Schedule has been successfully created and posted")

}

func DeleteSchedule(w http.ResponseWriter, r *http.Request) {

	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	deleteSchedule(id)

	fmt.Println("Schedule for station " + id + " has been deleted")
}

func main() {

	http.HandleFunc("/getInfo/", GetStation)
	http.HandleFunc("/newStation", PostStation)
	http.HandleFunc("/newSchedule", PostSchedule)
	http.HandleFunc("/delete/", DeleteSchedule)
	http.ListenAndServe(":8080", nil)

}
