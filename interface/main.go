package main

import (
	"fmt"
)

type Person struct {
	name string
}

func (p Person) GetName() string {
	return p.name
}

//Pointer receiver is required to change the struct value
func (p *Person) SetName(name string) {
	p.name = name
}

type Employee struct {
	Person
	role string
}

type Getter interface {
	GetName() string
}

func GetName(m Getter) {
	fmt.Printf("Hello %s \n", m.GetName())
}

type Setter interface {
	SetName(name string)
}

func SetName(s Setter, name string) {
	s.SetName(name)
}

func main() {
	//person can be value of type Person or pointer to type Person (&Person).
	//In this example person is value of type Person
	person := Person{
		name: "Jan Haans",
	}

	GetName(person)
	//pointer because SetName method of Person has pointer receiver
	SetName(&person, "Valentina Haans")
	GetName(person)

	employee := Employee{
		Person: Person{
			name: "AM Gobati",
		},
		role: "baas",
	}

	GetName(employee)
	SetName(&employee, "Dino Gobati")
	GetName(employee)

}
