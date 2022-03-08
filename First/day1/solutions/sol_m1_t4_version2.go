package main

import "fmt"

func main() {

	var number int
	fmt.Println("Введите число")
	fmt.Scanf("%d\n", &number)

	var first int = number % 10
	var second int = number / 10 % 10
	var third int = number / 100

	if number%100 == number {
		fmt.Printf("Результат %d\n", (first + second))
	} else if number/100 < 10 {
		var result int = 1
		if first > 0 {
			result *= first // result = result * first
		}
		if second > 0 {
			result *= second // result = result * second
		}
		result *= third
		fmt.Printf("Результат %d\n", result)
	} else {
		fmt.Println("Операция не может быть выполнена")
	}
}
