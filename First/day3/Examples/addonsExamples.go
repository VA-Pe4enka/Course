package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	a = 6
	b = 12
	c := 5.0
	fmt.Println("Before", a, b, c)
	// Mеняем значения переменных
	a, b = b, a
	a, d := int(c), a
	fmt.Println("After", a, b, c, d)

	// Так обычно
	//var array [3]int
	// Так нельзя в Го, необходимо определить размер на этапе компиляции
	//var array1 [a]int

	if math.Sqrt(2) == math.Sqrt(2)*0.1*10 {
		fmt.Println("Все понятно")
	} else {
		fmt.Println("Не все понятно")
	}
	//123 = 1 * Pow(10, 2) + 2 * Pow(10, 1) + 3 * Pow(10, 0)
	/* Побитовые операции &(bit and), |(bin or), ^(bit xor), >>, <<
	Пример использования bit xor
	101 = 1 * Pow(2, 2) + 0 * Pow(2, 1) + 1 * Pow(2, 0)  == 5
	011 = 0 * Pow(2, 2) + 1 * Pow(2, 1) + 1 * Pow(2, 0)  == 3
	110 = 1 * Pow(2, 2) + 1 * Pow(2, 1) + 0 * Pow(2, 0)  == 6
	*/
	fmt.Println("5 ^ 3 = ", 5^3) // Out: 6
	fmt.Println("4 ^ 7 = ", 4^7) // Out: 3

	// Досрочное прерывание цикла с помощью break
	number := 1
	for number < 10 {
		fmt.Println(number)
		if number == 4 {
			break
		}
		number++
	}
	fmt.Println("<for> have stopped.")
}
