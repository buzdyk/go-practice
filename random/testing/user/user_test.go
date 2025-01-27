package myuser

import "testing"

func Test_FullName(t *testing.T) {
	tests := []struct {
		name     string
		fname    string
		lname    string
		expected string
	}{
		{name: "simple test", fname: "John", lname: "Doe", expected: "John Doe"},
		{name: "only first name", fname: "John", expected: "John"},
		{name: "only last name", lname: "Doe", expected: "mr. Doe"},
		{name: "unknown name", expected: "Anonymous"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := User{FirstName: test.fname, LastName: test.lname}
			if v := u.FullName(); v != test.expected {
				t.Errorf("Fser.FullName() = %v, expected %v", v, test.expected)
			}
		})
	}
}
