package positional_line_test

import (
	"fmt"
	"testing"

	"github.com/vert-capital/positional_line"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{
			[]struct {
				Field1 string  `positional:"10"`
				Field2 float64 `positional:"5,nofloat,leftpad"`
			}{
				{"hello", 123},
				{"456", 45},
			},
			"hello     12300\n456        4500",
		},
		{
			struct {
				Field3 float64 `positional:"15"`
			}{
				123.70,
			},
			"123.70         ",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Marshal(%v)", test.input), func(t *testing.T) {
			actual, err := positional_line.Marshal(test.input)

			if err != nil {
				t.Errorf("Marshal(%v) raised error %v", test.input, err)
			}

			if actual != test.expected {
				t.Errorf("Marshal(%v) = %v, expected %v", test.input, actual, test.expected)
			}
		})
	}
}
