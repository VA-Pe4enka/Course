package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//resOne, _ := http.Get("http://golang.org")

	//Простой пользовательский HTTP-клиент
	client := &http.Client{Timeout: time.Second}
	resOne, err := client.Get("http://golang.org/doc")
	if err != nil {
		log.Fatal(err)
	}
	//Свойство res.Status содержит текстовое сообщение о состоянии
	//Свойство res.StatusCode содержит код ошибки в виде целого числа.
	// GET = 200, POST = 201, PUT/PATCH = 202, DELETE = 204

	fmt.Println(resOne.Status, resOne.StatusCode)

	body, _ := ioutil.ReadAll(resOne.Body)
	resOne.Body.Close()
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(body)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	resTwo, err := http.Get("http://example.com/test.zip")
	if err != nil {
		fmt.Println("A error occured")
	}
	fmt.Println(resTwo.Status, resTwo.StatusCode)
}
