package main

import (
	"fmt"
	"sort"
)

func status(name string){
	grades := map[string]float64{"Petr": 0, "Anna": 5}
	grade := grades[name]
	if grade < 3 {
		fmt.Printf("%s не прошел(а) экзамен!\n", name)
	}
}

func statusImprove(name string){
	grades := map[string]float64{"Petr": 0, "Anna": 5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("нет оценок для %s.\n", name)
	} else if grade < 3 {
		fmt.Printf("%s не прошел(а) экзамен!\n", name)
	}
}

func main(){
	status("Petr")
	status("Ivan")

	statusImprove("Petr")
	statusImprove("Ivan")

	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool

	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)

	// Проверка ключа(есть ли ключ в словаре)
	_, ok = counters["a"]
	fmt.Println(ok)
	_, ok = counters["c"]
	fmt.Println(ok)



	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	// //delete key(pair)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	gradesNew := map[string]float64{"iFirst": 74.2, "fSecond": 86.5, "bThird": 59.7, "fFirst": 59.7}
	var names []string
	// 	Итерация по ключам
	for name := range gradesNew {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s имеет оценки за выступление %0.1f%%\n", name, gradesNew[name])
	}
	// 	Итерация по ключам и значениям
	for k, v := range gradesNew {
		fmt.Printf("Key: %v, value: %v\n", k, v)
	}
}