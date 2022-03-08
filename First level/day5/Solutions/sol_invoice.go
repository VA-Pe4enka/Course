package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"com.additions/utils"
)

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

func validateName(value string) bool {
	match, ok := regexp.MatchString(`^.{1,100}$`, value)
	return ok == nil && match
}

func validateCount(value string) bool {
	result, err := strconv.Atoi(value)
	return err == nil && result > 0
}

func validateBuyerName(value string) bool {
	match, ok := regexp.MatchString(`^([\w\pL]+\s?){2,}$`, value)
	return ok == nil && match
}

func validatePhone(value string) bool {
	match, ok := regexp.MatchString(`^\+?\d{10,11}$`, value)
	return ok == nil && match
}

func validateAddress(value string) bool {
	match, ok := regexp.MatchString(`([1-9]\d{5}), (([\w\d\pL\s.]+,){3,4}) ([\w\d\pL\s.]+)`, value)
	return ok == nil && match
}

type Product struct {
	name      string
	count     int
	buyerName string
	phone     string
	address   string
}

type Products interface {
	Name() string
	SetName()
	Count() int
	SetCount()
	BuyerName() string
	SetBuyerName()
	Phone() string
	SetPhone()
	Address() string
	SetAddress(value string)
}

func (p Product) Name() string {
	return p.name
}

func (p *Product) SetName() {
	var value string
	for !validateName(value) {
		value = strings.Trim(utils.ReadLine("Введите имя продукта"), " \r\n")
	}
	p.name = value
}

func (p Product) Count() int {
	return p.count
}

func (p *Product) SetCount() {
	var value string
	for !validateCount(value) {
		value = strings.Trim(utils.ReadLine("Введите количество"), " \r\n")
	}
	p.count, _ = strconv.Atoi(value)
}

func (p Product) BuyerName() string {
	return p.buyerName
}

func (p *Product) SetBuyerName() {
	var value string
	for !validateBuyerName(value) {
		value = strings.Trim(utils.ReadLine("Введите имя покупателя"), " \r\n")
	}
	p.buyerName = value
}

func (p Product) Phone() string {
	return p.phone
}

func (p *Product) SetPhone() {
	var value string
	for !validatePhone(value) {
		value = strings.Trim(utils.ReadLine("Введите телефон покупателя"), " \r\n")
	}
	p.phone = value
}

func (p Product) Address() string {
	return p.address
}

func (p *Product) SetAddress(value string) {
	var input string
	if len(value) > 0 && validateAddress(value) {
		fmt.Println("Введите адрес доставки (оставьте пустым, чтобы подставить указанный)")
		input = strings.Trim(utils.ReadLine(fmt.Sprintf("[%s]", value)), " \r\n")
		if len(input) == 0 {
			input = value
		}
	}
	for !validateAddress(input) {
		input = strings.Trim(utils.ReadLine("Введите адрес доставки"), " \r\n")
	}
	p.address = input
}

func NewProduct() Products {
	return &Product{}
}

func printTable(iterable []Products) {
	fmt.Printf("Наименование\tКоличество\tПокупатель\tКонтактный тел.\tАдрес доставки\n")
	for _, product := range iterable {
		fmt.Printf("%s\t%d\t%s\t%s\t%s\n",
			product.Name(),
			product.Count(),
			product.BuyerName(),
			product.Phone(),
			product.Address())
	}
}

func main() {
	var (
		count   int
		goods   []Products
		address string
	)

	utils.Ask("Какое количество товаров в накладной?", "%d", &count)

	goods = make([]Products, 0, count)

	for i := 0; i < count; i++ {
		fmt.Printf("Информация для товара #%d\n", i+1)
		product := NewProduct()

		product.SetName()
		product.SetCount()
		product.SetBuyerName()
		product.SetPhone()
		product.SetAddress(address)

		address = product.Address()

		goods = append(goods, product)
	}

	printTable(goods)

}
