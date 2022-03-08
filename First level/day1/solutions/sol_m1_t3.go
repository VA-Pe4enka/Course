/*
Посчитать размер вклада при определенном количестве лет и проценте.
Использовать ежегодную капитализацию.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var years int
	var sum float64
	var percent float64

	fmt.Println("Введите сумму вклада:")
	fmt.Scanf("%f\n", &sum)
	fmt.Println("Введите процент:")
	fmt.Scanf("%f\n", &percent)
	fmt.Println("На сколько лет?")
	fmt.Scanf("%d\n", &years)
	// Общая сумма начисленных процентов
	//total := math.Pow(1 + percent/100.0, float64(years))
	//fmt.Printf("total: %v\n", total)
	// math.Pow() - возведение в степень.
	result := sum * math.Pow(1 + percent/100.0, float64(years))

	fmt.Printf("Через %d лет будет %.2f\n", years, result)

}
