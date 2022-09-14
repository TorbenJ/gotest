package person

import "fmt"

type Person interface {
	Name() string
	FirstName() string
	LastName() string
	Age() int
}

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p *person) Age() int {
	return p.age
}

func (p *person) FirstName() string {
	return p.firstName
}

func (p *person) LastName() string {
	return p.lastName
}

func New(firstName string, lastName string, age int) Person {
	return &person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func (p *person) Name() string {
	return fmt.Sprintf("%v %v", p.firstName, p.lastName)
}
