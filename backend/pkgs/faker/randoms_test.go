package faker

import (
	"testing"
)

const Loops = 500

func ValidateUnique(values []string) bool {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] == values[j] {
				return false
			}
		}
	}
	return true
}

func Test_GetRandomString(t *testing.T) {
	t.Parallel()
	// Test that the function returns a string of the correct length
	generated := make([]string, Loops)

	faker := NewFaker()

	for i := 0; i < Loops; i++ {
		generated[i] = faker.Str(10)
	}

	if !ValidateUnique(generated) {
		t.Error("Generated values are not unique")
	}
}

func Test_GetRandomEmail(t *testing.T) {
	t.Parallel()
	// Test that the function returns a string of the correct length
	generated := make([]string, Loops)

	faker := NewFaker()

	for i := 0; i < Loops; i++ {
		generated[i] = faker.Email()
	}

	if !ValidateUnique(generated) {
		t.Error("Generated values are not unique")
	}
}

func Test_GetRandomBool(t *testing.T) {
	t.Parallel()

	trues := 0
	falses := 0

	faker := NewFaker()

	for i := 0; i < Loops; i++ {
		if faker.Bool() {
			trues++
		} else {
			falses++
		}
	}

	if trues == 0 || falses == 0 {
		t.Error("Generated boolean don't appear random")
	}
}

func Test_RandomNumber(t *testing.T) {
	t.Parallel()

	f := NewFaker()

	const MIN = 0
	const MAX = 100

	last := MIN - 1

	for i := 0; i < Loops; i++ {
		n := f.Num(MIN, MAX)

		if n == last {
			t.Errorf("RandomNumber() failed to generate unique number")
		}

		if n < MIN || n > MAX {
			t.Errorf("RandomNumber() failed to generate a number between %v and %v", MIN, MAX)
		}
	}
}
