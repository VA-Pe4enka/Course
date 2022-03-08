package main

import "fmt"

type SomeError string

// type error interface {
// 	Error()
// }

func (s SomeError) Error() string {
	return string(s)
}
func main() {
	err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
	fmt.Println(err.Error())
	fmt.Println(err)

	var error1 error
	error1 = SomeError("Just a message")
	qty, error2 := fmt.Println(error1)
	fmt.Println(qty)
	fmt.Println(error2)
	fmt.Println(error1.Error())
}
