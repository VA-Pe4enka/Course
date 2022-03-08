package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address // это тип встроенной структуры
}

type Employee struct {
	Name   string
	Salary float64
	Address
}

type Address struct {
	PostalCode string
	State      string
	City       string
	Street     string

}