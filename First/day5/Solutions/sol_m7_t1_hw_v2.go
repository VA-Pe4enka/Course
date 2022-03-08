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
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type order struct {
	product
	customer
}

type product struct {
	productName   string
	productNumber int
}

type customer struct {
	customerName  string
	customerPhone int
	address
}

type address struct {
	postalCode int
	city       string
	street     string
	house      string
	apartment  string
}

func (o order) askString() (str string) {
	str = readAnswer()
	return
}

func (o order) askIntAndStr() (i int, str string, err error) {
	str = readAnswer()
	i, err = strconv.Atoi(str)
	return
}

func (o order) askInt() (i int, err error) {
	str := readAnswer()
	i, err = strconv.Atoi(str)
	return i, err
}

func main() {
	var order order
	var isPostalCode, isAddress, isPhone, isProductNameNumber bool

	fmt.Println("Введите индекс(6 цифр)")
	for {
		PostalCodeInt, PostalCode, err := order.askIntAndStr()
		if len(PostalCode) != 6 && err == nil {
			fmt.Println("Длина не 6 символов")
		} else if len(PostalCode) == 6 && err != nil {
			fmt.Println("Пожалуйста, используйте цифры")
		} else if len(PostalCode) != 6 && err != nil {
			fmt.Println("Ты что издеваешься? Это похоже на индекс?")
		} else if len(PostalCode) == 6 && err == nil {
			isPostalCode = true
			order.postalCode = PostalCodeInt
		}
		if isPostalCode {
			break
		}
	}
	fmt.Println("Введите адрес, через запятую (Город, Улица, Дом, Квартира)\nПожалуйста, не оставляйте поля пустыми, если нет улицы или номера дома, на из месте поставьте \"1\"")
	for {
		address := strings.Split(order.askString(), ", ")

		if len(address) != 4 {
			fmt.Println("Не, ты ты чего? По человечески прошу, вот пример (Город, Улица, Дом, Квартира)\nА то до утра продолжать будем\nПожалуйста, не оставляйте поля пустыми, если нет улицы или номера дома, на их месте поставьте \"1\"")
		} else {
			isAddress = true
			order.city = address[0]
			order.street = address[1]
			order.house = address[2]
			order.apartment = address[3]
		}

		if isAddress {
			break
		}
	}
	fmt.Println("как Вас зовут?")
	order.customerName = order.askString()
	fmt.Println("Ваш номер телефона? Используйте формат 81231231212. Не должен начинаться с 0")
	for {
		phoneInt, phone, err := order.askIntAndStr()

		if len(phone) != 11 && err == nil {
			fmt.Println("Длина не 11 символов")
		} else if len(strconv.Itoa(phoneInt)) == 10 && err == nil {
			fmt.Println("Ну нормально разговаривали, что начинаешь? давай без нолей вначале")
		} else if len(phone) == 11 && err != nil {
			fmt.Println("Пожалуйста, используйте цифры")
		} else if len(phone) != 11 && err != nil {
			fmt.Println("Ты что издеваешься? Это похоже на номер телефона?")
		} else if len(phone) == 11 && err == nil {
			isPhone = true
			order.customerPhone = phoneInt
		}
		if isPhone {
			break
		}
	}

	fmt.Println("Что покупаете?")
	order.productName = order.askString()
	fmt.Printf("Сколько единиц товара: %v?\n", order.productName)
	for {
		numberInt, err := order.askInt()
		if numberInt == 0 && err == nil {
			fmt.Println("Если тебе ничего не надо, то может и заказывать не стоит?")
		} else if err != nil {
			fmt.Println("Пожалуйста, используйте цифры")
		} else if numberInt != 0 && err == nil {
			isProductNameNumber = true
			order.productNumber = numberInt
		}
		if isProductNameNumber {
			break
		}
	}

	fmt.Printf("Добрый день, %v\n\tВы заказали %v единиц товара: %v\n\tДоставка по адресу: \n\t\tИндекс: %v, город %v, улица %v, дом №%v, квартира №%v \n\t\tВаш контактный телефон: %v",
		order.customerName, order.productNumber, order.productName, order.postalCode, order.city, order.street, order.house, order.apartment, order.customerPhone)
}

func readAnswer() string {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data := strings.TrimSpace(line)
	return data
}
