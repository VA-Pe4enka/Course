package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type twoDigits struct {
	FirstNum  int `json:"firstNumber"`
	SecondNum int `json:"secondNumber"`
	Sum       int `json:"summa"`
}

var digits twoDigits

func main() {
	rand.Seed(time.Now().Unix())
	digits.FirstNum = rand.Intn(100)
	digits.SecondNum = rand.Intn(100)
	digits.Sum = digits.FirstNum + digits.SecondNum
	out, err := json.MarshalIndent(digits, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
	http.HandleFunc("/first", FirstHandler)
	http.HandleFunc("/second", SecondHandler)
	http.HandleFunc("/summa", SumHandler)
	http.ListenAndServe(":1234", nil)
}

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(digits.FirstNum)
}
func SecondHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(digits.SecondNum)
}
func SumHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(digits.Sum)
}
