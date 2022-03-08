package main

import (
	"calendar"
	"fmt"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("День рождения")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

	err = event.SetTitle("Слишком длинное значение для Title")
	if err != nil {
		log.Fatal(err)
	}
}
