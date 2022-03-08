package main

import (
	"fmt"
)

func main() {
	var n int
	var oldN int
	fmt.Println("Введите число: ")
	fmt.Scanf("%d\n", &n)
	oldN = n
	a := n % 10
	n /= 10
	b := n % 10
	n /= 10
	c := n % 10
	d := n / 10
	n = 1000*a + 100*b + 10*c + d

	if oldN == n {
		fmt.Printf("%d - палиндром\n", oldN)
	} else {
		fmt.Printf("%d - не палиндром\n", oldN)
	}
}
