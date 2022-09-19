package MongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"homework/Models"
	"log"
	"os"
)

var EmployeeCollection *mongo.Collection
var CTX = context.TODO()

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AddEmployee() {

	stack := []Models.Stack{}
	stack = append(stack, Models.Stack{"Golang"})
	stack = append(stack, Models.Stack{"Python"})
	stack = append(stack, Models.Stack{"django"})

	projs := []Models.Projects{}
	projs = append(projs, Models.Projects{"Ozon", "SeniorDev"})
	projs = append(projs, Models.Projects{"WB", "TeamLead"})
	projs = append(projs, Models.Projects{"VK", "SeniorDev"})

	desWork := Models.DesiredWork{"Google"}
	desStack := Models.DesiredStack{"Golang"}

	employee := Models.Employee{
		Name:         "Jhon",
		Stack:        stack,
		Projects:     projs,
		DesiredWork:  desWork,
		DesiredStack: desStack,
	}

	stack = nil
	projs = nil

	insertResult, err := EmployeeCollection.InsertOne(CTX, employee)
	Check(err)
	fmt.Println("First employee has been added.", insertResult)

	stack = append(stack, Models.Stack{"JS"})
	stack = append(stack, Models.Stack{"Vue3"})
	stack = append(stack, Models.Stack{"Python"})

	projs = append(projs, Models.Projects{"Telegram", "TeamLead"})
	projs = append(projs, Models.Projects{"Inst", "TeamLead"})
	projs = append(projs, Models.Projects{"VK", "TeamLead"})

	desWork = Models.DesiredWork{"Yandex"}
	desStack = Models.DesiredStack{"Python"}

	employee = Models.Employee{
		Name:         "Alice",
		Stack:        stack,
		Projects:     projs,
		DesiredWork:  desWork,
		DesiredStack: desStack,
	}

	stack = nil
	projs = nil

	insertResult, err = EmployeeCollection.InsertOne(CTX, employee)
	Check(err)
	fmt.Println("Second employee has been added.", insertResult)

	stack = append(stack, Models.Stack{"Java"})
	stack = append(stack, Models.Stack{"Golang"})
	stack = append(stack, Models.Stack{"Python"})

	projs = append(projs, Models.Projects{"Ozon", "TeamLead"})
	projs = append(projs, Models.Projects{"Telegram", "TeamLead"})
	projs = append(projs, Models.Projects{"WB", "TeamLead"})

	desWork = Models.DesiredWork{"Google"}
	desStack = Models.DesiredStack{"Python"}

	employee = Models.Employee{
		Name:         "Timmy",
		Stack:        stack,
		Projects:     projs,
		DesiredWork:  desWork,
		DesiredStack: desStack,
	}

	stack = nil
	projs = nil

	insertResult, err = EmployeeCollection.InsertOne(CTX, employee)
	Check(err)
	fmt.Println("Third employee has been added.", insertResult)

	stack = append(stack, Models.Stack{"JS"})
	stack = append(stack, Models.Stack{"Vue3"})
	stack = append(stack, Models.Stack{"React"})

	projs = append(projs, Models.Projects{"VK", "TeamLead"})
	projs = append(projs, Models.Projects{"Yandex", "TeamLead"})
	projs = append(projs, Models.Projects{"Uber", "TeamLead"})

	desWork = Models.DesiredWork{"Telegram"}
	desStack = Models.DesiredStack{"Java"}

	employee = Models.Employee{
		Name:         "Michael",
		Stack:        stack,
		Projects:     projs,
		DesiredWork:  desWork,
		DesiredStack: desStack,
	}

	stack = nil
	projs = nil

	insertResult, err = EmployeeCollection.InsertOne(CTX, employee)
	Check(err)
	fmt.Println("Forth employee has been added.", insertResult)

	stack = append(stack, Models.Stack{"Java"})
	stack = append(stack, Models.Stack{"Python"})
	stack = append(stack, Models.Stack{"django"})

	projs = append(projs, Models.Projects{"Inst", "TeamLead"})
	projs = append(projs, Models.Projects{"Telegram", "TeamLead"})
	projs = append(projs, Models.Projects{"WB", "TeamLead"})

	desWork = Models.DesiredWork{"Yandex"}
	desStack = Models.DesiredStack{"Python"}

	employee = Models.Employee{
		Name:         "Caitlin",
		Stack:        stack,
		Projects:     projs,
		DesiredWork:  desWork,
		DesiredStack: desStack,
	}

	stack = nil
	projs = nil

	insertResult, err = EmployeeCollection.InsertOne(CTX, employee)
	Check(err)
	fmt.Println("Fifth employee has been added.", insertResult)

}

func Choose() int {
	var choose int
	fmt.Print("Выберите позицию: ")
	fmt.Fscan(os.Stdin, &choose)
	//проверка на ввод только числа и рекурсия функции в случае неправильного ввода

	return choose
}

func FindEmployee() {

	fmt.Println("Выберите фильтр: ")
	fmt.Println("1. Стек технологий")
	fmt.Println("2. Проекты над которыми работали")
	fmt.Println("3. Желаемая позиция")
	fmt.Println("4. Желаемый стек")
	fmt.Println()
	fmt.Println("0. Выход")

}
