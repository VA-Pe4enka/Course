package main

import (
	"example.com/input"
	"fmt"
	"log"
)

func main() {
	number, err := input.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(number)
}
