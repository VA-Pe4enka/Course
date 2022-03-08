package main

import "fmt"

type Number int

//Создание метода Add у типа Number
func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n) + otherNumber)
}
//Создание метода Subtract у типа Number
func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n) - otherNumber)
}

func (n *Number) AddChange(otherNumber int) {
	*n += Number(otherNumber)
}

func main() {
	ten := Number(10)
	ten.Add(5)
	ten.Subtract(4)
	fmt.Println("Before:", ten)
	ten.AddChange(10)
	fmt.Println("After:", ten)

	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}