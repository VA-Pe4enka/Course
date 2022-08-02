package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type nums struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func FirstHandl(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(100)
	first := []nums{
		nums{"first", num},
	}
	json.NewEncoder(w).Encode(first)
}

func SecondHandl(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(100)
	second := []nums{
		nums{"second", num},
	}
	json.NewEncoder(w).Encode(second)
}

func SummaHandl(w http.ResponseWriter, r *http.Request) {
	result := rand.Intn(100) + rand.Intn(100)
	summa := []nums{
		nums{"summa", result},
	}
	json.NewEncoder(w).Encode(summa)
}

func main() {
	http.HandleFunc("/first", FirstHandl)
	http.HandleFunc("/second", SecondHandl)
	http.HandleFunc("/summa", SummaHandl)
	http.ListenAndServe(":1234", nil)
}
