/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/
package main

import (
	"fmt"
	"strconv"
)

// Частичная проверка
func examination(x1, y1, x2, y2, x3, y3 int) bool {
	var flg bool
	if x1 == x2 && x2 == x3 {
		flg = false
	} else if y1 == y2 && y2 == y3 {
		flg = false
	} else {
		flg = true
	}
	return flg
}

func Square(x1, y1, x2, y2, x3, y3 int) float64 {
	var number int
	var S float64
	n := (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)
	nn := fmt.Sprintf("%v", n)
	if nn[0] == '-' {
		nnn := nn[1:]
		ss, err := strconv.Atoi(nnn)
		if err != nil {
			fmt.Println(err)
		}
		number = ss
		S = 0.5 * float64(number)
	} else {
		number = n
		S = 0.5 * float64(number)
	}

	return S
}

func rightTriangle(x1, y1, x2, y2, x3, y3 int) bool {
	var flg bool
	ab := (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
	bc := (x3-x2)*(x3-x2) + (y3-y2)*(y2-y1)
	ac := (x3-x1)*(x3-x1) + (y3-y1)*(y2-y1)

	d1 := bc + ac
	d2 := ab + ac
	d3 := ab + bc
	if ab == d1 || bc == d2 || ac == d3 {
		flg = true
	} else {
		flg = false
	}
	return flg
}

func main() {
	var x1, y1, x2, y2, x3, y3 int
	fmt.Println("Введете координаты")
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	fmt.Scan(&x3)
	fmt.Scan(&y3)
	result := examination(x1, y1, x2, y2, x3, y3)
	if result {
		square := Square(x1, y1, x2, y2, x3, y3)
		fmt.Println("площадь", square)
		right := rightTriangle(x1, y1, x2, y2, x3, y3)
		if right {
			fmt.Println("Треугольник прямоугольный")
		} else {
			fmt.Println("Треугольник не прямоугольный")
		}
	} else {
		fmt.Println("По данным координатам нельзя построить треугольник")
	}
}
