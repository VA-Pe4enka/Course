package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!")
}
func deutschHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hallo, web!")
}

func main() {

	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/hallo", deutschHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
