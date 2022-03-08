package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL string
	Size int
}

func responseSize(url string, channel chan Page){
	fmt.Println("Getting", url)
	response , err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://yandex.ru",
					 "https://1c.ru",
					 "https://google.com",
					 "https://golang.com/doc",
					 "https://ibm.com",
					}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i< len(urls); i++ {
		page := <- pages
		fmt.Printf("%s: %d bytes\n", page.URL, page.Size)
	}

}