package server

import (
	"fmt"
	"homework/homework/info"
	"homework/homework/models"
	"html/template"
	"log"
	"net/http"
)

type Service struct{}

type Server interface {
	Handler()
}

type ViewData struct {
	models.Person
}

var testTemplate *template.Template

func (service *Service) Handler() {
	var err error
	testTemplate, err = template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", getInfo)
	fmt.Println("Server is going on!")
	http.ListenAndServe(":8080", nil)

}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	info.AddPerson()

	person := info.Person

	vd := ViewData{person}

	err := testTemplate.Execute(w, vd)
	if err != nil {
		log.Fatal(err)
	}
}
