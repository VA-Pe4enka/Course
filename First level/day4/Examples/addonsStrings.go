package main

import "fmt"

func foo() (first, second int) {
	first = 1
	second = first + 5
	return
}

func main() {
	s1 := "hello"
	s2 := "Привет"
	fmt.Println(len(s1), len(s2))

	slice := []string{"do", "re", "mi", "fa", "sol"}
	array := [5]string{"do", "re", "mi", "fa", "sol"}
	fmt.Println(slice[3:], array[3:])

	var1, var2, var3 := 512, "hello", true
	fmt.Println(var1, var2, var3)

	fmt.Println(foo())
}