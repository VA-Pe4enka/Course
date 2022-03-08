package main

import (
	"fmt"
	"strings"
)

type Product struct {
	ProductName string
	Count       string
	Person      Person
}

type Person struct {
	Name       string
	Surname    string
	Patronymic string
	Phone      string
	Address    Address
}

type Address struct {
	Index  string
	City   string
	Street string
	House  string
	Flat   string
}

func main() {

setProdName:
	product := Product{}
	fmt.Print("Enter the product name: ")
	fmt.Scanf("%s\n", &product.ProductName)
	if len(product.ProductName) < 1 || len(product.ProductName) > 100 {
		fmt.Println("Invalid product name!!! It`s shouldn`t be over 100 and less 1 symbols.")
		product.ProductName = ""
		goto setProdName
	}
	if strings.ContainsAny(product.ProductName, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.ProductName = ""
		goto setProdName
	}

setCount:
	fmt.Print("Enter count: ")
	fmt.Scanf("%s\n", &product.Count)
	if strings.ContainsAny(product.Count, "abcdefghiklmnopqrstvxyz") || strings.ContainsAny(product.Count, "ABCDEFGHIKLMNOPQRSTVXYZ") {
		fmt.Println("Invalid count! It`s should contain only digits.")
		product.Count = ""
		goto setCount
	}
	if len(product.Count) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Count = ""
		goto setCount
	}
	if strings.ContainsAny(product.Count, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Count = ""
		goto setCount
	}
	if len(product.Count) == 1 && strings.ContainsAny(product.Count, "0") {
		fmt.Println("Invalid value! It`s shouldn`t be 0")
		product.Count = ""
		goto setCount
	}

setUserName:
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s\n", &product.Person.Name)
	if strings.ContainsAny(product.Person.Name, "0123456789") {
		fmt.Println("Invalid name! It`s shouldn`t contain digits.")
		product.Person.Name = ""
		goto setUserName
	}
	if len(product.Person.Name) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Person.Name = ""
		goto setUserName
	}
	if strings.ContainsAny(product.Person.Name, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Person.Name = ""
		goto setUserName
	}

setSurname:
	fmt.Print("Enter your surname: ")
	fmt.Scanf("%s\n", &product.Person.Surname)
	if strings.ContainsAny(product.Person.Surname, "0123456789") {
		fmt.Println("Invalid surname! It`s shouldn`t contain digits.")
		product.Person.Surname = ""
		goto setSurname
	}
	if len(product.Person.Surname) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Person.Surname = ""
		goto setSurname
	}
	if strings.ContainsAny(product.Person.Surname, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Person.Surname = ""
		goto setSurname
	}

setPatronymic:
	fmt.Print("Enter your patronymic: ")
	fmt.Scanf("%s\n", &product.Person.Patronymic)
	if strings.ContainsAny(product.Person.Patronymic, "0123456789") {
		fmt.Println("Invalid patronymic! It`s shouldn`t contain digits.")
		product.Person.Patronymic = ""
		goto setPatronymic
	}
	if len(product.Person.Patronymic) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Person.Patronymic = ""
		goto setPatronymic
	}
	if strings.ContainsAny(product.Person.Patronymic, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Person.Patronymic = ""
		goto setPatronymic
	}

setPhone:
	fmt.Print("Enter your phone number: ")
	fmt.Scanf("%s\n", &product.Person.Phone)
	if strings.ContainsAny(product.Person.Phone, "abcdefghiklmnopqrstvxyz") ||
		strings.ContainsAny(product.Person.Phone, "ABCDEFGHIKLMNOPQRSTVXYZ") {
		fmt.Println("Invalid phone number!  It`s should contain only digits.")
		product.Person.Phone = ""
		goto setPhone
	}
	if len(product.Person.Phone) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Person.Phone = ""
		goto setPhone
	}
	if strings.ContainsAny(product.Person.Phone, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Person.Phone = ""
		goto setPhone
	}

setIndex:
	fmt.Print("Enter your index: ")
	fmt.Scanf("%s\n", &product.Person.Address.Index)
	if len(product.Person.Address.Index) > 6 {
		fmt.Println("Invalid index! It`s shouldn`t be over 6 numbers.")
		product.Person.Address.Index = ""
		goto setIndex
	}
	if len(product.Person.Address.Index) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Person.Address.Index = ""
		goto setIndex
	}
	if strings.ContainsAny(product.Person.Address.Index, "abcdefghiklmnopqrstvxyz") ||
		strings.ContainsAny(product.Person.Address.Index, "ABCDEFGHIKLMNOPQRSTVXYZ") {
		fmt.Println("Invalid index!  It`s should contain only digits.")
		product.Person.Address.Index = ""
		goto setIndex
	}
	if strings.ContainsAny(product.Person.Address.Index, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Person.Address.Index = ""
		goto setIndex
	}

setCity:
	fmt.Print("Enter your city: ")
	fmt.Scanf("%s\n", &product.Person.Address.City)
	if strings.ContainsAny(product.Person.Address.City, "0123456789") {
		fmt.Println("Invalid city! It`s shouldn`t contain digits.")
		product.Person.Address.City = ""
		goto setCity
	}
	if len(product.Person.Address.City) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Person.Address.City = ""
		goto setCity
	}
	if strings.ContainsAny(product.Person.Address.City, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Person.Address.City = ""
		goto setCity
	}

setStreet:
	fmt.Print("Enter your street: ")
	fmt.Scanf("%s\n", &product.Person.Address.Street)
	if strings.ContainsAny(product.Person.Address.Street, "0123456789") {
		fmt.Println("Invalid street!  It`s shouldn`t contain digits.")
		product.Person.Address.Street = ""
		goto setStreet
	}
	if len(product.Person.Address.Street) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Person.Address.Street = ""
		goto setStreet
	}
	if strings.ContainsAny(product.Person.Address.Street, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Person.Address.Street = ""
		goto setStreet
	}

setHouse:
	fmt.Print("Enter your house: ")
	fmt.Scanf("%s\n", &product.Person.Address.House)
	if len(product.Person.Address.House) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Person.Address.House = ""
		goto setHouse
	}
	if strings.ContainsAny(product.Person.Address.House, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Person.Address.House = ""
		goto setHouse
	}

setFlat:
	fmt.Print("Enter your flat: ")
	fmt.Scanf("%s\n", &product.Person.Address.Flat)
	if strings.ContainsAny(product.Person.Address.Flat, "abcdefghiklmnopqrstvxyz") ||
		strings.ContainsAny(product.Person.Address.Flat, "ABCDEFGHIKLMNOPQRSTVXYZ") {
		fmt.Println("Invalid flat!  It`s shouldn`t contain digits.")
		product.Person.Address.Flat = ""
		goto setFlat
	}
	if len(product.Person.Address.Flat) == 0 {
		fmt.Println("Folder shouldn`t be empty!")
		product.Person.Address.Flat = ""
		goto setFlat
	}
	if strings.ContainsAny(product.Person.Address.Flat, "_!@#$%^&") {
		fmt.Println("Invalid value! It`s shouldn`t contain any of _!@#$%^&")
		product.Person.Address.Flat = ""
		goto setFlat
	}

	fmt.Println()
	fmt.Println("__________________________________________")
	fmt.Println()
	fmt.Println("Product name: ", product.ProductName)
	fmt.Println("Count: ", product.Count)
	fmt.Println()
	fmt.Println("Your name: ", product.Person.Name, " ", product.Person.Surname, " ", product.Person.Patronymic)
	fmt.Println("Your phone number: ", product.Person.Phone)
	fmt.Println()
	fmt.Println("Your address: ", product.Person.Address.Index, " ", product.Person.Address.City, " ", product.Person.Address.Street, " ", product.Person.Address.House, " ", product.Person.Address.Flat)
	fmt.Println()
	fmt.Println("__________________________________________")
	fmt.Println()
}
