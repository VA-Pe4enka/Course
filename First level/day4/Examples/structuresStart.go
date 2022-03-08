package main

import "fmt"

type Person struct {
	name	string
	age		int
	adult	bool
}

func changeName(p *Person) {
	p.name = "other name"
}

func main() {
	//name  string
	//age   int
	//adult bool
	var person struct {
		name	string
		age		int
		adult	bool
	}
	person.name = "first"
	person.age = 18
	person.adult = true
	fmt.Println(person.name)

	person1 := Person{name: "first", age: 12, adult: false}
	person2 := Person{name: "second", age: 20, adult: true}

	fmt.Println(person1)
	fmt.Println(person2)
	changeName(&person1)
	fmt.Println(person1)
}