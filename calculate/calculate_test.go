package calculate

import (
	"testing"
)

func TestFizz(t *testing.T) {
	const num int = 3

	expect := true
	actual := Fizz(num)

	if expect != actual {
		t.Errorf("expect: %t, but actual: %t", expect, actual)
	}
}

func TestBuzz(t *testing.T) {
	const num int = 5

	expect := true
	actual := Buzz(num)

	if expect != actual {
		t.Errorf("expect: %t, but actual: %t", expect, actual)
	}
}

func TestFizzBuzz(t *testing.T) {
	const num int = 15

	expect := true
	actual := FizzBuzz(num)

	if expect != actual {
		t.Errorf("expect: %t, but actual: %t", expect, actual)
	}
}
