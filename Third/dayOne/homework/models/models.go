package models

type Person struct {
	Name           FullName
	Education      Education
	WorkExperience []Company
	Projects       []Project
	Stack          []Stack
}

type FullName struct {
	Name       string
	SurName    string
	Patronymic string
}

type Company struct {
	Name string
}

type Education struct {
	Main          string
	Qualification []Courses
}

type Courses struct {
	Name string
}

type Project struct {
	Name string
}

type Stack struct {
	Name string
}
