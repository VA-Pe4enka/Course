package main

import "fmt"

func main() {
	var v1 int8 = 2
	var v11 uint = 6
	var v2 int16 = 3

	var v3 float32 = 3.3
	var v4 float64 = 1.1

	type Example int

	var e1 Example
	var e2 Example

	fmt.Println(int16(v1) + v2)
	fmt.Println(v3 + float32(v4))
	fmt.Println(int16(v3) + int16(v4))
	fmt.Println(v11)
	fmt.Println(e1 + e2)




}