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
