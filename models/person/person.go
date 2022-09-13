package person

import "fmt"

type Person interface {
	Name() string
}

type person struct {
	firstName string
	lastName  string
}

func New(firstName string, lastName string) Person {
	return &person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p *person) Name() string {
	return fmt.Sprintf("%v %v", p.firstName, p.lastName)
}
