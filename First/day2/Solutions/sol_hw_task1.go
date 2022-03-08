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
	const fuelPrice = 48.0

	var distance, fuelRate, totalPrice float64

	fmt.Println("Введите расстояние в км (50 - 10000):")
	fmt.Scanf("%f\n", &distance)
	fmt.Println("Введите расход топлива в литрах на 100 км (5-25 литров):")
	fmt.Scanf("%f\n", &fuelRate)

	totalPrice = fuelRate / 100.0 * distance * fuelPrice

	fmt.Printf("Итоговая сумма: %.2f\n", totalPrice)
}
