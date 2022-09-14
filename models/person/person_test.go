package person

import "testing"

func TestName(t *testing.T) {
	person := person{
		firstName: "John",
		lastName:  "Doe",
	}

	expected := "John Doe"
	actual := person.Name()

	assertEquals(t, "Name()", expected, actual)
}

func TestFirstName(t *testing.T) {
	person := person{firstName: "John"}

	expected := "John"
	actual := person.FirstName()

	assertEquals(t, "FirstName()", expected, actual)
}

func TestLastName(t *testing.T) {
	person := person{lastName: "Doe"}

	expected := "Doe"
	actual := person.LastName()

	assertEquals(t, "LastName()", expected, actual)
}

func TestAge(t *testing.T) {
	person := person{age: 25}

	expected := 25
	actual := person.Age()

	assertEquals(t, "Age()", expected, actual)
}

func assertEquals(t *testing.T, description string, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Fatalf("%v returned '%v', expected '%v'", description, actual, expected)
	}
}
