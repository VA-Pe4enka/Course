package main

import (
	"fmt"
	"example.com/input"
	"log"
)

func main() {
	number, err := input.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(number)
}