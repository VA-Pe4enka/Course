/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/
package main

import (
	"fmt"
)

func main() {
	var distance float64
	var consumption float64
	const (
		price float64 = 48.00
	)

	fmt.Println("Введите расстояние(50 - 10000 км):")
	fmt.Scanf("%f\n", &distance)
	if distance < 50 || distance > 10000 {
		fmt.Println("Неверно задано расстояние")
	} else {
		fmt.Println("Введите расход в литрах (5-25 литров) на 100 км:")
		fmt.Scanf("%f\n", &consumption)
		if consumption < 5 || consumption > 25 {
			fmt.Println("Неверно задан расход")
		} else {
			fmt.Printf("Cтоимость поездки в рублях: %.2f", 0.01*distance*consumption*price)
		}
	}

}
