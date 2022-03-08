/*
Написать 3 функции.
Вход:
1.Ширина участка
2.Длина участка
3.Стоимость кв.м. земли
4.Стоимость погонного метра забора

Дан участок земли прямоугольной формы.
Нужно посчитать стоимость самой земли и забора по периметру.

Первая функция принимает в качестве аргументов длину и ширину участка.
Возвращает периметр и площадь.

Вторая функция принимает в качестве аргументов периметр и стоимость погонного метра забора.
Возвращает итоговую сумму.

Третья функция принимает в качестве аргументов площадь и стоимость одного квадратного метра земли.
Возвращает итоговую сумму, используя коэффицент 1.3(умножить) .
*/
package main

import "fmt"

func main() {
	var width, height int
	var landPrice, fencePrice float64
	fmt.Println("Введите ширину участка : ")
	fmt.Scanf("%d\n", &width)
	fmt.Println("Введите длину участка : ")
	fmt.Scanf("%d\n", &height)
	fmt.Println("Введите стоимость кв.м. земли : ")
	fmt.Scanf("%f\n", &landPrice)
	fmt.Println("Введите стоимость п.м. забора : ")
	fmt.Scanf("%f\n", &fencePrice)
	perimetr, square := calculatePerimAndSquare(width, height)
	fmt.Printf("Периметр участка = %d п.м., площадь= %d кв.м\n", perimetr, square)

	fmt.Printf("Стоимость п.м. забора = %.2f руб.\n", calculateTotalFencePrice(perimetr, fencePrice))
	fmt.Printf("Стоимость кв.м. земли = %.2f руб.\n", calculateTotalLandPrice(square, landPrice))
}

func calculatePerimAndSquare(w, h int) (int, int) {
	return (h + w) * 2, h * w
}

func calculateTotalLandPrice(square int, price float64) float64 {
	return float64(square) * 1.3 * price
}

func calculateTotalFencePrice(perim int, price float64) float64 {
	return float64(perim) * price
}
