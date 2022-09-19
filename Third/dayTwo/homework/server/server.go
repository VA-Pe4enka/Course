package server

import (
	"fmt"
	"homework/Models"
	"homework/info"
	"html/template"
	"log"
	"net/http"
)

type Service struct{}

type Server interface {
	Handler()
}

type ViewData struct {
	Models.Employee
}

var testTemplate *template.Template

func (service *Service) Handler() {
	var err error
	testTemplate, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	//MongoDB.AddEmployee()
	//MongoDB.Choose()

	http.HandleFunc("/1", getInfo)

	fmt.Println("Server is going on!")
	http.ListenAndServe(":8080", nil)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	info.AddEmployee()

	employee := info.Employee

	vd := ViewData{employee}

	err := testTemplate.Execute(w, vd)
	{
		if err != nil {
			log.Fatal(err)
		}
	}
}
