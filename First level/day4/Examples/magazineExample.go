package main

import (
	"fmt"
	"magazine"
)

func main() {
	subscriber := magazine.Subscriber{
		Name:   "Petr Ivanov",
		Rate:   4.99,
		Active: true}

	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active:", subscriber.Active)

	var employee magazine.Employee
	employee.Name = "Student Second"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	var address magazine.Address
	address.Street = "Mira 1"
	address.City = "Moscow"
	address.State = "Msk"
	address.PostalCode = "123127"
	fmt.Println(address)

	subscriber2 := magazine.Subscriber{Name: "Student First"}
	subscriber2.Street = "Prospekt Lenina 111"
	subscriber2.City = "Pushkino"
	subscriber2.State = "Moscow region"
	subscriber2.PostalCode = "98765"
	fmt.Println("Street:", subscriber2.Street)
	fmt.Println("City:", subscriber2.City)
	fmt.Println("State:", subscriber2.State)
	fmt.Println("PostalCode:", subscriber2.PostalCode)
	fmt.Println(subscriber2)
	subscriber3 := magazine.Subscriber{Name: "Student Second"}
	fmt.Println(subscriber3)

}
