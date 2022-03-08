package main

import "fmt"

func main() {
	var (
		integer int
		result  string = ""
		// Integer version
		// result int
	)
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &integer)

	
	// String
	result = fmt.Sprintf("%s%d", result, integer % 10)
	result = fmt.Sprintf("%s%d", result, integer / 10 % 10)
	result = fmt.Sprintf("%s%d", result, integer / 100)

	// Integer
	//result += (integer % 10) *100
	//result += ((integer/10) % 10) * 10
	//result += integer / 100

	fmt.Printf("Результат: %s\n", result)

}