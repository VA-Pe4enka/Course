package info

import (
	"bufio"
	"fmt"
	"homework/homework/models"
	"log"
	"os"
	"strings"
)

func readLine(message string) string {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s> ", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}
	input = strings.Trim(strings.Trim(input, "\n"), "\r")

	return input
}

var Person models.Person

func FillName() models.FullName {

	Person.Name.Name = readLine("Укажите имя ")
	//проверка ввода(только буквы)
	Person.Name.SurName = readLine("Укажите фамилию ")
	//проверка ввода(только буквы)
	Person.Name.Patronymic = readLine("Укажите отчество ")
	//проверка ввода(только буквы)

	name := models.FullName{
		Name:       Person.Name.Name,
		SurName:    Person.Name.SurName,
		Patronymic: Person.Name.Patronymic,
	}

	return name
}
func FillEducation() models.Education {

	course := []models.Courses{}

	Person.Education.Main = readLine("Укажите основное образование ")
	//проверка ввода(нет спец. символов)

	for {
		course = append(course, FillCourse(readLine("Укажите название курса ")))
		//проверка ввода(нет спец. символов)

		f := readLine("Добавить еще один курс (y - да, n - нет)?")
		//проверка ввода(только y или n)
		if f == "n" {
			break
		}
	}

	education := models.Education{
		Main:          Person.Education.Main,
		Qualification: course,
	}

	return education

}

func FillCourse(name string) models.Courses {

	course := models.Courses{
		Name: name,
	}

	return course
}

func FillWork() []models.Company {
	work := []models.Company{}

	for {
		work = append(work, models.Company{readLine("Укажите место работы ")})
		//проверка ввода(нет спец. символов)

		f := readLine("Добавить еще одно место работы (y - да, n - нет)?")
		//проверка ввода(только y или n)
		if f == "n" {
			break
		}
	}

	return work
}

func FillProjs() []models.Project {

	projects := []models.Project{}

	for {
		projects = append(projects, models.Project{readLine("Укажите ваши проекты ")})
		//проверка ввода(нет спец. символов)

		f := readLine("Добавить еще один проект (y - да, n - нет)?")
		//проверка ввода(только y или n)
		if f == "n" {
			break
		}
	}

	return projects
}

func FillStack() []models.Stack {
	stack := []models.Stack{}

	for {
		stack = append(stack, models.Stack{readLine("Укажите ваш стек ")})

		f := readLine("Добавить еще одну технологию (y - да, n - нет)?")
		//проверка ввода(только y или n)
		if f == "n" {
			break
		}
	}

	return stack
}

func AddPerson() {
	person := models.Person{
		Name:           FillName(),
		Education:      FillEducation(),
		WorkExperience: FillWork(),
		Projects:       FillProjs(),
		Stack:          FillStack(),
	}

	Person = person

	fmt.Println(person)
}
