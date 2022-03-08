package main

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (my MyType) MethodWithoutParameters() {
	panic("implement me")
}

// func (m MyType) MethodWithoutParameters() {
// 	fmt.Println("MethodWithoutParameters called")
// }
func (my MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}
func (my MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}
func (my MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
}

func main() {
	var value MyInterface
	value = MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
