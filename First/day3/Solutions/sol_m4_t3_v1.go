/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var x1, y1, x2, y2, x3, y3 int
	//var str string
	fmt.Println("Введите координаты трех точек (x1, y1, x2, y2, x3, y3) : ")
	//Version 1
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	numbers := strings.Fields(line)
	//fmt.Println(data)

	// Version 2
	//fmt.Scanf("%s\n", &str)
	// strings.Split() разделяет строку по словам, используя пробел в качестве разделителя.
	//numbers:= strings.Split(str, ",")
	// fmt.Printf("%v, %#T", numbers,numbers)
	x1, err = strconv.Atoi(numbers[0])
	check(err)
	y1, err = strconv.Atoi(numbers[1])
	check(err)
	x2, _ = strconv.Atoi(numbers[2])
	y2, _ = strconv.Atoi(numbers[3])
	x3, _ = strconv.Atoi(numbers[4])
	y3, _ = strconv.Atoi(numbers[5])

	fmt.Printf("Координаты точек(x1, y1, x2, y2, x3, y3) - (%d, %d, %d, %d, %d, %d)\n", x1, y1, x2, y2, x3, y3)
	if checkExistenceTriangle(x1, y1, x2, y2, x3, y3) {
		fmt.Println("По заданным точкам можно построить треугольник.")
		fmt.Printf("Площадь треугольника = %.2f кв.м.\n", squareTriangle(x1, y1, x2, y2, x3, y3))
		if checkIsRightTriangle(x1, y1, x2, y2, x3, y3) {
			fmt.Println("Треугольник является прямоугольным.")
		} else {
			fmt.Println("Треугольник не прямоугольный.")
		}

	} else {
		fmt.Println("Треугольник по заданным точкам построить нельзя.\n")
	}
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Pow(math.Pow((float64)(x2-x1), 2)+math.Pow((float64)(y2-y1), 2), 0.5)
}

func checkExistenceTriangle(x1, y1, x2, y2, x3, y3 int) bool {
	var edge1Distance, edge2Distance, edge3Distance float64
	edge1Distance = distance(x1, y1, x2, y2)
	edge2Distance = distance(x2, y2, x3, y3)
	edge3Distance = distance(x1, y1, x3, y3)
	return (edge1Distance+edge2Distance > edge3Distance &&
		edge1Distance+edge3Distance > edge2Distance &&
		edge2Distance+edge3Distance > edge1Distance)
}

func squareTriangle(x1, y1, x2, y2, x3, y3 int) float64 {
	var edge1Distance, edge2Distance, edge3Distance, semiPerim float64
	edge1Distance = distance(x1, y1, x2, y2)
	edge2Distance = distance(x2, y2, x3, y3)
	edge3Distance = distance(x1, y1, x3, y3)
	semiPerim = (edge1Distance + edge2Distance + edge3Distance) * 0.5
	return math.Pow(semiPerim*(semiPerim-edge1Distance)*(semiPerim-edge2Distance)*(semiPerim-edge3Distance), 0.5)
}

func checkIsRightTriangle(x1, y1, x2, y2, x3, y3 int) bool {
	var edge1Distance, edge2Distance, edge3Distance float64
	edge1Distance = distance(x1, y1, x2, y2)
	edge2Distance = distance(x2, y2, x3, y3)
	edge3Distance = distance(x1, y1, x3, y3)
	return math.Abs(edge1Distance*edge1Distance+edge2Distance*edge2Distance-edge3Distance*edge3Distance) <= 0.01 ||
		math.Abs(edge1Distance*edge1Distance+edge3Distance*edge3Distance-edge2Distance*edge2Distance) <= 0.01 ||
		math.Abs(edge3Distance*edge3Distance+edge2Distance*edge2Distance-edge1Distance*edge1Distance) <= 0.01
}
