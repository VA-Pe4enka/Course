/*
## Задача № 1
Написать API для указанных маршрутов(endpoints)
"/first"  // Случайное число
"/second" // Случайное число
"/summa"  // Сумма двух случайных чисел

результат вернуть в виде JSON
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()
	rand.Seed(5)
	a, b := rand.Intn(50), rand.Intn(25)
	sum := a + b
	numbers := map[string]int{"First": a, "Second": b, "Sum": sum}
	numbers2, err := json.MarshalIndent(numbers, "", "")
	fmt.Println(string(numbers2))
	SoftErr(err)

	newMux.HandleFunc("/first", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "<font size=\"72\"><em><strong>%d</strong></em></font>", a)
		SoftErr(err)
	})
	newMux.HandleFunc("/second", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "<font size=\"72\"><em><strong>%d</strong></em></font>", b)
		SoftErr(err)
	})
	newMux.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "<font size=\"72\"><em><strong>%d + %d = %d</strong></em></font>", a, b, sum)
		SoftErr(err)
	})
	err = http.ListenAndServe(":1235", newMux)
	SoftErr(err)
}

func SoftErr(err error) {
	if err != nil {
		log.Fatalln("wrong data")
	}
}
