package main

import "fmt"

type mytype []string

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type point struct {
	x, y int
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func distance(p1 point, p2 point) float64 {
	return float64(p1.x) + float64(p2.x)
}

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	var auto car
	auto.name = "Машина"
	auto.topSpeed = 323
	fmt.Println("Name:", auto.name)
	fmt.Println("Top speed:", auto.topSpeed)

	var bolts part
	bolts.description = "Болт 16"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)

	var subscriber1 subscriber
	subscriber1.name = "Петр Иванов"
	fmt.Println("Name:", subscriber1.name)
	var subscriber2 subscriber
	subscriber2.name = "Иван Петров"
	fmt.Println("Name:", subscriber2.name)

	showInfo(bolts)

	partResult := minimumOrder("Болт 16")
	fmt.Println(partResult.description, partResult.count)

	var s1 mytype
	s1 = append(s1, "hello")
	fmt.Println(s1)

	type pointLocal struct {
		x, y int
	}
	p1 := pointLocal{x: 1, y: 2}
	fmt.Println(p1)

	p2 := point{x: 10, y: 20}
	p3 := point{x: 1, y: 2}
	fmt.Println(distance(p2, p3))
}
