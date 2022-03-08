package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var bottles, check int
var str string

var x1, x2, x3, y1, y2, y3, side1, side2, side3 float64

var size int

func countSides(x1, x2, x3, y1, y2, y3 float64) {
	side1 = math.Sqrt((math.Pow(x1-x2, 2)) + (math.Pow(y1-y2, 2)))
	side2 = math.Sqrt((math.Pow(x1-x3, 2)) + (math.Pow(y1-y3, 2)))
	side3 = math.Sqrt((math.Pow(x2-x3, 2)) + (math.Pow(y2-y3, 2)))
}

func m4t2(bottles int) string {
	if bottles < 0 || bottles > 200 {
		fmt.Println("Enter correct value!!!")
		os.Exit(1)
	}

	for i := 1; i <= 190; i += 10 {
		if bottles == i {
			if bottles == 11 {
				str = strconv.Itoa(bottles) + " " + "бутылок"
			} else {
				str = strconv.Itoa(bottles) + " " + "бутылка"
			}

		}
	}

	if bottles < 10 {
		check = bottles
	}
	if bottles >= 10 && bottles < 200 {
		check = bottles % 10
	}

	switch check {
	case 2, 3, 4:
		if bottles % 100 == 12 || bottles % 100 == 13 || bottles % 100 == 14 {
			str = strconv.Itoa(bottles) + " " + "бутылок"
		} else {
			str = strconv.Itoa(bottles) + " " + "бутылки"
		}
	case 5, 6, 7, 8, 9:
		str = strconv.Itoa(bottles) + " " + "бутылок"
	}

	return str
}

func m4t3T1(x1, x2, x3, y1, y2, y3 float64) bool {
	var check bool
	var side1, side2, side3 float64

	side1 = math.Sqrt((math.Pow(x1-x2, 2)) + (math.Pow(y1-y2, 2)))
	side2 = math.Sqrt((math.Pow(x1-x3, 2)) + (math.Pow(y1-y3, 2)))
	side3 = math.Sqrt((math.Pow(x2-x3, 2)) + (math.Pow(y2-y3, 2)))

	if ((side1 + side2) > side3) && ((side1 + side3) > side2) && ((side2 + side3) > side1) {
		check = true
	} else {
		check = false
	}

	return check

}

func m4t3T2(x1, x2, x3, y1, y2, y3 float64) float64 {
	var p float64

	countSides(x1, x2, x3, y1, y2, y3)
	p = (side1 + side2 + side3) / 2

	return math.Sqrt(p * (p - side1) * (p - side2) * (p - side3))

}

func m4t3T3(x1, x2, x3, y1, y2, y3 float64) bool {
	var check bool

	countSides(x1, x2, x3, y1, y2, y3)

	if (math.Pow(side1, 2)+math.Pow(side2, 2) == math.Pow(side3, 2)) ||
		(math.Pow(side1, 2)+math.Pow(side3, 2) == math.Pow(side2, 2)) ||
		(math.Pow(side3, 2)+math.Pow(side2, 2) == math.Pow(side1, 2)) {
		check = true
	} else {
		check = false
	}

	return check
}

func m4t4(size int) {
	if size < 0 || size > 20 {
		fmt.Println("Enter correct value!!!")
		os.Exit(1)
	}

	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board[i][j] = 0
			} else {
				board[i][j] = 1
			}
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%v ", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Enter count of bottles: ")
	fmt.Scanf("%d\n", &bottles)
	fmt.Println(m4t2(bottles))

	fmt.Println("Enter the cords: ")
	fmt.Print("x1= ")
	fmt.Scanf("%f\n", &x1)
	fmt.Print("x2= ")
	fmt.Scanf("%f\n", &x2)
	fmt.Print("x3= ")
	fmt.Scanf("%f\n", &x3)
	fmt.Print("y1= ")
	fmt.Scanf("%f\n", &y1)
	fmt.Print("y2= ")
	fmt.Scanf("%f\n", &y2)
	fmt.Print("y3= ")
	fmt.Scanf("%f\n", &y3)

	fmt.Println("Does a triangle exist? ", m4t3T1(x1, x2, x3, y1, y2, y3))
	fmt.Println("S= ", m4t3T2(x1, x2, x3, y1, y2, y3))
	fmt.Println("Is it right triangle? ", m4t3T3(x1, x2, x3, y1, y2, y3))

	fmt.Println("Enter board size: ")
	fmt.Scanf("%d\n", &size)

	m4t4(size)
}
