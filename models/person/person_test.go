package person

import "testing"

func TestName(t *testing.T) {
	person := person{
		firstName: "John",
		lastName:  "Doe",
	}

	expected := "John Doe"
	actual := person.Name()

	if expected != actual {
		t.Fatalf("Name() returned '%v', expected '%v'", actual, expected)
	}
}

func TestFirstName(t *testing.T) {
	person := person{firstName: "John"}

	expected := "John"
	actual := person.FirstName()

	if expected != actual {
		t.Fatalf("FirstName() returned '%v', expected '%v'", actual, expected)
	}
}

func TestLastName(t *testing.T) {
	person := person{lastName: "Doe"}

	expected := "Doe"
	actual := person.LastName()

	if expected != actual {
		t.Fatalf("LastName() returned '%v', expected '%v'", actual, expected)
	}
}
