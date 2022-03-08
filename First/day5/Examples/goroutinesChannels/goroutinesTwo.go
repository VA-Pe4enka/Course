package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url+":", len(body), "bytes.")
}

func main() {
	go responseSize("https://yandex.ru")
	go responseSize("https://1c.ru")
	go responseSize("https://google.com")
	go responseSize("https://golang.com")
	go responseSize("https://ibm.com")
	time.Sleep(2 * time.Second)
}
