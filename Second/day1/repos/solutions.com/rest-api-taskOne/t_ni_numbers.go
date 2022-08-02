package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Number struct {
	Name   string `json:"Name"`
	Number int    `json:"Number"`
}

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/first", FirstHandler)
	http.HandleFunc("/second", SecondHandler)
	http.HandleFunc("/summa", SummaHandler)
	http.ListenAndServe(":1234", nil)
}

func number() int {
	return rand.Intn(100)
}

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	num1 := number()
	result := Number{"first", num1}

	json.NewEncoder(w).Encode(result)
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {
	num2 := number()
	result := Number{"second", num2}

	json.NewEncoder(w).Encode(result)
}

func SummaHandler(w http.ResponseWriter, r *http.Request) {
	num1 := number()
	num2 := number()
	sum := num1 + num2
	result := Number{"summa", sum}

	json.NewEncoder(w).Encode(result)
}
