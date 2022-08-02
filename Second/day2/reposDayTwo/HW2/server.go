package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Person struct {
	ID         string `json:"id"`
	FIO        string `json:"fio"`
	DEPARTMENT string `json:"department"`
	POSITION   string `json:"position"`
	WORKTIME   int    `json:"worktime"`
}

type Note struct {
	ID        string    `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Result    int       `json:"result"`
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	employeeData := getEmployee(id)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(employeeData)

}

func PostEmployee(w http.ResponseWriter, r *http.Request) {

	postEmployee()

	fmt.Println("New user has been successfully created")

}

func GetAllTime(w http.ResponseWriter, r *http.Request) {

}

func start(w http.ResponseWriter, r *http.Request) {
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	createNote(id)
}

func end(w http.ResponseWriter, r *http.Request) {
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	updateNote(id)
	updateEmployee(id)
}

func main() {

	http.HandleFunc("/get/", GetEmployee)
	http.HandleFunc("/post", PostEmployee)
	http.HandleFunc("/getAll", GetAllTime)
	http.HandleFunc("/start/", start)
	http.HandleFunc("/end/", end)

	http.ListenAndServe(":8080", nil)

}
