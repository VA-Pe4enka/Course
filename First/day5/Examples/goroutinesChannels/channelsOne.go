package main

import (
	"fmt"
	"time"
)

func greeting(myChannel chan string) {
	// Записываем строку в канал
	myChannel <- "hi"
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	// считываем из канала
	result := <-myChannel
	fmt.Println(result)

	go send(myChannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
