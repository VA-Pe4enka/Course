package main

import (
	"fmt"
)

func main() {
	var ranks map[string]int

	ranks = make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["tree"])
	fmt.Println(ranks)
	fmt.Printf("%#v\n", ranks)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])

	// Создание мапы со значениями
	ranksNew := map[string]int{"gold": 1, "silver": 2, "bronze": 3}
	fmt.Println(ranksNew["gold"])
	fmt.Println(ranksNew["silver"])

	//multiline map
	elementsNew := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
	}
	fmt.Println(elementsNew["H"])
	fmt.Println(elementsNew["Li"])

	numbers := make(map[string]int)
	numbers["first"] = 12
	fmt.Printf("%#v\n", numbers["first"])
	fmt.Printf("%#v\n", numbers["second"])

	words := make(map[string]string)
	words["english"] = "hello"
	fmt.Printf("%#v\n", words["english"])
	fmt.Printf("%#v\n", words["spanich"])
	fmt.Println(words)

	// Можно сразу выполнить операцию инкремента
	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])

	var nilMap map[int]string
	fmt.Printf("%#v, %v\n", nilMap, len(nilMap))
	//panic error -> " assignment to entry in nil map"
	//nilMap[3] = "three"

	var myMap map[int]string = make(map[int]string)
	myMap[3] = "three"
	fmt.Printf("%#v\n", myMap)

}
