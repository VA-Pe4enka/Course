/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

реализовать несколько методов у типа "Накладная"

createReader == NewReader
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	CustomerName string
	Items        []Item
	Phone        string
	Address
}

type Item struct {
	Name   string
	Amount int
}

type Address struct {
	PostalCode string
	City       string
	Street     string
	House      string
	Apartment  string
}

func showInfo(o Order) {
	fmt.Println("Накладная")

	fmt.Println("ФИО:", o.CustomerName)
	fmt.Println("Телефон:", o.Phone)

	fmt.Println("Адрес:", o.PostalCode, o.City, o.Street, o.House, o.Apartment)

	fmt.Println("Заказ:")
	fmt.Println(" № | Наименование | Количество")

	for index, item := range o.Items {
		fmt.Printf("%2v | %12v | %2v\n", index+1, item.Name, item.Amount)
	}

}

func makeOrder(customerName string, phone string, address Address) (Order, error) {
	var o Order
	if customerName == "" {
		return o, errors.New("empty field: <customerName>")
	} else {
		o.CustomerName = customerName
	}
	o.Phone = phone
	o.PostalCode = address.PostalCode
	o.City = address.City
	o.Street = address.Street
	o.House = address.House
	o.Apartment = address.Apartment
	o.Items = make([]Item, 0)

	return o, nil
}

func (o *Order) addItem(name string, amount int) {
	var item Item
	item.Name = name
	item.Amount = amount

	o.Items = append(o.Items, item)
}

func readLine(message string) string {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s> ", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}
	input = strings.Trim(strings.Trim(input, "\n"), "\r")

	return input
}

func main() {
	customerName := readLine("Введите ФИО")
	phone := readLine("Введите номер телефона")
	var addr Address

	fmt.Println("Введите адрес:")
	addr.PostalCode = readLine("Почтовый индекс")
	addr.City = readLine("Город")
	addr.Street = readLine("Улица")
	addr.House = readLine("Дом")
	addr.Apartment = readLine("Квартира")

	order, err := makeOrder(customerName, phone, addr)
	if err != nil {
		panic(err)
	}
	flag := true
	i := 1

	for flag {
		fmt.Println("Товар", i, ":")
		i++
		name := readLine("Наименование")
		amount, _ := strconv.Atoi(readLine("Количество"))

		order.addItem(name, amount)

		f := readLine("Добавить еще один товар (y - да, n - нет)?")
		flag = f == "y"
	}

	fmt.Println()

	showInfo(order)

}
