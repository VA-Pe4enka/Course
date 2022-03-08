package main

import "fmt"

func main() {
	var nameStudent string
	var number int
	var numberFloat float64
	nameStudent = "First"
	var second string = "Second line"
	third := "Third line"
	third = "Change value"
	var value bool

	value = true
	fmt.Println("First line")
	fmt.Println("Hello, Golang")
	fmt.Println(nameStudent)
	fmt.Println(number)
	fmt.Println(numberFloat)
	fmt.Println(second)
	fmt.Printf("%v, %T\n", third, third)
	fmt.Printf("%v, %T\n", value, value)
}
