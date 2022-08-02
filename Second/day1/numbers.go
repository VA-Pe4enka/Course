package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Numbers struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func FirstNumHandler(w http.ResponseWriter, r *http.Request) {

	first := rand.Intn(100)
	value := []Numbers{
		Numbers{"first", first},
	}
	json.NewEncoder(w).Encode(value)

}

func SecondNumHandler(w http.ResponseWriter, r *http.Request) {

	second := rand.Intn(100)
	value := []Numbers{
		Numbers{"second", second},
	}
	json.NewEncoder(w).Encode(value)

}

func SummHandler(w http.ResponseWriter, r *http.Request) {

	summ := rand.Intn(100) + rand.Intn(100)
	value := []Numbers{
		Numbers{"summ", summ},
	}
	json.NewEncoder(w).Encode(value)

}

func main() {

	rand.Seed(time.Now().Unix())

	http.HandleFunc("/first", FirstNumHandler)
	http.HandleFunc("/second", SecondNumHandler)
	http.HandleFunc("/summ", SummHandler)
	http.ListenAndServe(":8080", nil)

}
