package pad_test

import (
	"testing"

	"github.com/vert-capital/positional_line/pad"
)

func TestLeft(t *testing.T) {
	tests := []struct {
		str         string
		size        int
		pad         string
		expected    string
		expectedErr bool
	}{
		{"hello", 10, "*", "*****hello", false},
		{"world", 5, "*", "world", false},
		{"world", 6, "*", "*world", false},
		{"", 5, "*", "*****", false},
		{"longer string", 1, "*", "l", false},
		{"teste error", -1, "*", "", true},
		{"teste error empty", 1, "", "", true},
	}

	for _, test := range tests {
		result, err := pad.Left(test.str, test.size, test.pad)
		if result != test.expected {
			t.Errorf("Left(%q, %d, %q) = %q; want %q", test.str, test.size, test.pad, result, test.expected)
		}
		if (err != nil) != test.expectedErr {
			t.Errorf("Left(%q, %d, %q) returned error %v; want error %t", test.str, test.size, test.pad, err, test.expectedErr)
		}
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		str         string
		size        int
		pad         string
		expected    string
		expectedErr bool
	}{
		{"hello", 10, "*", "hello*****", false},
		{"world", 5, "*", "world", false},
		{"world", 6, "*", "world*", false},
		{"", 5, "*", "*****", false},
		{"longer string", 1, "*", "l", false},
		{"teste error", -1, "*", "", true},
		{"teste error empty", 1, "", "", true},
	}

	for _, test := range tests {
		result, err := pad.Right(test.str, test.size, test.pad)
		if result != test.expected {
			t.Errorf("Left(%q, %d, %q) = %q; want %q", test.str, test.size, test.pad, result, test.expected)
		}
		if (err != nil) != test.expectedErr {
			t.Errorf("Left(%q, %d, %q) returned error %v; want error %t", test.str, test.size, test.pad, err, test.expectedErr)
		}
	}
}
