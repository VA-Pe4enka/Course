/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/
package main

import "fmt"

func main() {
	var distance float64
	fmt.Println("Введите расстояние в км (50 - 10000 км) : ")
	fmt.Scanf("%f\n", &distance)
	var consumption float64
	fmt.Println("Введите расход в литрах на 100 км (5-25 литров) : ")
	fmt.Scanf("%f\n", &consumption)
	const price = 48

	if 50 <= distance && distance <= 10000 && 5 <= consumption && consumption <= 25 {
		fmt.Printf("Стоимость поездки составит %.2f руб.\n", distance/100.0*consumption*price)
	} else {
		fmt.Println("Введены неверные данные. Расчет не может быть выполнен.")
	}
}
