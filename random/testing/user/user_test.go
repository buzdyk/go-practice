package myuser

import "testing"
import "github.com/stretchr/testify/assert"

func Test_FullName(t *testing.T) {
	tests := []struct {
		name  string
		fname string
		lname string
		want  string
	}{
		{name: "simple test", fname: "John", lname: "Doe", want: "John Doe"},
		{name: "only first name", fname: "John", want: "John"},
		{name: "only last name", lname: "Doe", want: "mr. Doe"},
		{name: "unknown name", want: "Anonymous"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := User{FirstName: test.fname, LastName: test.lname}
			assert.Equal(t, u.FullName(), test.want)
		})
	}
}
