package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Бум-Бум!")
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.MakeSound()
	}
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything(Whistle("Канарейка"))
}
