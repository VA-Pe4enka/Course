package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	http.HandleFunc("/first", FirstHandler)
	http.HandleFunc("/second", SecondHandler)
	http.HandleFunc("/summa", SumHandler)
	http.ListenAndServe(":1234", nil)
}

type Summa struct {
	First  int
	Second int
	Result int
}

var s Summa // Все поля равны 0

func FirstHandler(w http.ResponseWriter, r *http.Request) {

	s.First = rand.Intn(100)

	json.NewEncoder(w).Encode(s.First)
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {

	s.Second = rand.Intn(100)

	json.NewEncoder(w).Encode(s.Second)
}

func SumHandler(w http.ResponseWriter, r *http.Request) {

	s.Result = s.First + s.Second

	json.NewEncoder(w).Encode(s)
}
