package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subscriber) {
	fmt.Println("Имя:", s.name)
	fmt.Println("Стоимость подписки:", s.rate)
	fmt.Println("Активный?", s.active)

}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func printInfoPointer(s *subscriber) {
	fmt.Println("Имя:", s.name)
	fmt.Println("Стоимость подписки:", s.rate)
	fmt.Println("Активный?", s.active)
}

func defaultSubsPointer(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s // возвращаем адрес нового подписчика
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
func main() {
	subscriber1 := defaultSubscriber("Иван Петров")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Сидр Иванов")
	printInfo(subscriber2)

	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)

	//Это уже не структура, а указатель на структуру
	subscriber3 := defaultSubsPointer("Алексей Михайлов")
	applyDiscount(subscriber3)
	printInfoPointer(subscriber3)

	subscriber4 := defaultSubsPointer("Сергей Алексеев")
	printInfoPointer(subscriber4)
}
