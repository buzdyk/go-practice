package abs

import (
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{name: "positive number", value: 22.0, expected: 22.0},
		{name: "negative number", value: -22.0, expected: 22.0},
		{name: "zero", value: float64(0), expected: float64(0)},
	}

	for _, test := range tests { // цикл по всем тестам
		t.Run(test.name, func(t *testing.T) {
			if v := Abs(test.value); v != test.expected {
				t.Errorf("Sum() = %v, expected %v", v, test.expected)
			}
		})
	}
}
