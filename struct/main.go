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

type Messenger interface {
	GetName() string
}

func SendMessage(m Messenger) {
	fmt.Printf("Het gaat goed met %s \n", m.GetName())
}

func main() {
	person := Person{
		name: "Jan Haans",
	}

	fmt.Printf("Hello %s \n", person.GetName())
	SendMessage(person)
	person.SetName("Valentina Haans")
	fmt.Printf("Hello %s \n", person.GetName())
	SendMessage(person)

	employee := Employee{
		Person: Person{
			name: "AM Gobati",
		},
		role: "baas",
	}

	fmt.Printf("Hello %s \n", employee.GetName())
	SendMessage(employee)
	employee.SetName("Dino Gobati")
	fmt.Printf("Hello %s \n", employee.GetName())
	SendMessage(employee)
}
