package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value int = 3
	var pointer *int = &value // &value - это адрес переменной value, а pointer - это указатель
	fmt.Println(pointer)
	fmt.Println(*pointer) // Через адрес получаем значение переменной, на которую указывает данный указатель

	var param myStruct
	param.myField = 3
	var pointer2 *myStruct = &param

	fmt.Println((*pointer2).myField) // явное получение доступа к самой структуре
	fmt.Println(pointer2.myField)    // неявное получение доступа к самой структуре

	pointer2.myField = 9
	fmt.Println(pointer2.myField)

}
