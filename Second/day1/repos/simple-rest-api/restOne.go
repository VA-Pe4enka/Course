package main

import (
	"encoding/json"
	"net/http"
)

type Animal struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
}

func AnimalsHandler(w http.ResponseWriter, r *http.Request) {
	animals := []Animal{
		Animal{"Alice", "Cat"},
		Animal{"Bob", "Cat"},
		Animal{"Spike", "Dog"},
	}

	json.NewEncoder(w).Encode(animals)
	// с тем же самым результатом (the same)
	// encoder := json.NewEncoder(w)
	// encoder.Encode(animals)
}

func main() {
	http.HandleFunc("/animals", AnimalsHandler)
	// Сервер будет слушать на всех сетевых интерфейсах
	// "0.0.0.0:1234" - только IPv4
	// ":1234" - IPv4 and IPv6
	http.ListenAndServe(":1234", nil)
}
