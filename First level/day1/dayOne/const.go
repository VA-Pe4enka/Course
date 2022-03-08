package main 

import (
	// Каждый модуль в отдельной строке
	"fmt"
	"math" 
) 
const (
	lang = "golang"
	ide = "vscode"
	version = 1.17
)
func main() { 
	//Множественная инициализация
	var i, j, k int = 1, 2, 3 

	const ( 
		// iota начинает с 0 и добавляет 1 к каждому значению.
		first = iota 
		second 
		third 
		fourth 
	) 


	var ( 
		name = "First" 
		age = 23 
		prof = "Student" 
	) 
			
	//Распечатать тип переменной
	fmt.Printf("i: %v, type: %T\n" , i, i) 

	//Проверим значения 
	fmt.Printf("i: %v, j: %v, k: %v\n", i, j, k) 

	fmt.Println("name: " , name) 
	fmt.Println("age: " , age) 
	fmt.Println("prof: " , prof) 

	fmt.Println("lang and ide:", lang + "/" + ide) 
	fmt.Println("version: ", version - 0.5) 
	fmt.Printf("math module: %4.3f\n", math.Sqrt(version))
	
	//version = version + 1 // Error: cannot assign to version (declared const)

	fmt.Println("first:" , first) 
	fmt.Println("second:" , second) 
	fmt.Println("third:" , third) 
	fmt.Println("fourth:" , fourth)

}