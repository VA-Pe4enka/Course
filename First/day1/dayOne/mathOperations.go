package main

import (
	"fmt"
	//	"strconv"
)

func main() {
	var result float64
	var d float64 = 10
	var e float64 = 4

	a := 42
	b := 20
	c := a / b
	result = d / e
	// Результат - это целое число(но тип float остался)
	var check float64 = 20 / 3

	fmt.Printf("type of c: %T\n", c)
	fmt.Printf("%v, type: %T\n", result, result)
	fmt.Printf("%v, type: %T\n", check, check)

	fmt.Print("Инкремент:")
	// check = check + 1
	check++
	fmt.Println("check =", check)
	fmt.Print("Декремент:")
	// c = c - 1
	c--
	fmt.Println("c =", c)

	// Операция взятия остатка
	newValue := int(check) % 3
	fmt.Println(newValue)

	// s1 := string(124)
	// fmt.Println(s1)
	// s2 := strconv.Itoa(124)
	// fmt.Println(s2[0], s2[1], s2[2])
	// fmt.Println(s2)
}
