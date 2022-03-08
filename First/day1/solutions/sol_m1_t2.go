// Задача №2
// Вход: трехзначное число.
// Найти первую, вторую и последнюю цифры заданного трехзначного числа.
package main

import "fmt"

func main() {
	var number int
	// Блок ввода
	fmt.Println("enter 3-digits number")
	fmt.Scanf("%d", &number)

	secondNumber := number - (number/100)*100 // number % 100
	// Получаем разряд сотен
	fmt.Printf("first: %d \n", number/100)

	// Получаем разряд десятков
	fmt.Printf("middle: %d \n", secondNumber/10) // (number / 10) % 10

	// Получаем разряд единиц
	fmt.Printf("last: %d \n", number%10)
}
