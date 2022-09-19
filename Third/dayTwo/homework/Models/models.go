package Models

type Employee struct {
	Name         string
	Stack        []Stack
	Projects     []Projects
	DesiredWork  DesiredWork
	DesiredStack DesiredStack
}

type Stack struct {
	Name string
}

type Projects struct {
	Name string
	Post string
}

type DesiredWork struct {
	Name string
}

type DesiredStack struct {
	Name string
}
