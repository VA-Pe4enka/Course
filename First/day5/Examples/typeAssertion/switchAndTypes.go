package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type MyType int

func awardPrize() {
	switch rand.Intn(5) {
	case 1:
		fmt.Println("Вы выиграли первый приз!")
	case 2:
		fmt.Println("Вы выиграли второй приз!")
	case 3:
		fmt.Println("Вы выиграли третий приз!")
	default:
		fmt.Println("увы... Ничего")
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v,", rand.Intn(10))
	}
	fmt.Println()
	awardPrize()
	var x interface{}

	// x = nil
	// x = 2
	// x = 4.1
	// x = true
	// x = "hello"
	x = MyType(23)

	// Пример утверждения типа(Type assertions)
	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x is", v)
	case bool, string:
		fmt.Println("x is bool or string")
	default:
		fmt.Println("type unknown")
	}
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
}
