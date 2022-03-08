package main

import (
	"fmt"
	"math"
	"log"
)

func main() {
// Вариант с max и min:
	var x, y, z float64

	fmt.Print("Three numbers: \n>>> ")
	_, err := fmt.Scanf("%f%f%f", &x, &y, &z)
	if err != nil {
		log.Fatal("Not a number")
	}

	min := math.Min(math.Min(x, y), z)
	max := math.Max(math.Max(x, y), z)
	mid := x + y + z - min - max

	fmt.Println(min, mid, max)
}