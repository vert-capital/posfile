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

func TestMarshalError(t *testing.T) {
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
			_, err := positional_line.Marshal(test.input)

			if err != nil {
				t.Errorf("Marshal(%v) raised error %v", test.input, err)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	line := "hello     12300"

	type TestStruct struct {
		Field1 string  `positional:"10"`
		Field2 float64 `positional:"5,nofloat,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != "hello" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, "hello")
	}

	if test.Field2 != 12300 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, 12300)
	}
}

func TestUnmarshalDifferentPositions(t *testing.T) {
	line := "12300     hello"

	type TestStruct struct {
		Field1 float64 `positional:"5,nofloat,leftpad"`
		Field2 string  `positional:"10"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != 12300 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, 12300)
	}

	if test.Field2 != "hello" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, "hello")
	}
}

func TestUnmarshalWithIntField(t *testing.T) {
	line := "hello     12345"

	type TestStruct struct {
		Field1 string `positional:"10"`
		Field2 int    `positional:"5,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != "hello" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, "hello")
	}

	if test.Field2 != 12345 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, 12345)
	}
}

// func TestUnmarshalWithShortLine(t *testing.T) {
// 	line := "short"

// 	type TestStruct struct {
// 		Field1 string `positional:"10"`
// 		Field2 int    `positional:"5,leftpad"`
// 	}

// 	var test TestStruct

// 	err := positional_line.Unmarshal(line, &test)

// 	if err == nil {
// 		t.Errorf("Expected error for short line, but got none")
// 	}
// }

func TestUnmarshalWithInvalidInt(t *testing.T) {
	line := "hello     abcde"

	type TestStruct struct {
		Field1 string `positional:"10"`
		Field2 int    `positional:"5,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err == nil {
		t.Errorf("Expected error for invalid int, but got none")
	}
}

func TestUnmarshalWithInvalidFloat(t *testing.T) {
	line := "hello     123.45"

	type TestStruct struct {
		Field1 string `positional:"10"`
		Field2 int    `positional:"5,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err == nil {
		t.Errorf("Expected error for invalid float, but got none")
	}
}

func TestUnmarshalWithInvalidSize(t *testing.T) {
	line := "hello     12345"

	type TestStruct struct {
		Field1 string `positional:"10"`
		Field2 int    `positional:"abc,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err == nil {
		t.Errorf("Expected error for invalid size, but got none")
	}
}

func TestUnmarshalWithValidFloat(t *testing.T) {
	line := "hello     123.45"

	type TestStruct struct {
		Field1 string  `positional:"10"`
		Field2 float64 `positional:"5,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != "hello" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, "hello")
	}

	if test.Field2 != 123.45 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, 123.45)
	}
}

func TestUnmarshalWithZeroFill(t *testing.T) {
	line := "hello     00123"

	type TestStruct struct {
		Field1 string  `positional:"10"`
		Field2 float64 `positional:"5,zerofill"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != "hello" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, "hello")
	}

	if test.Field2 != 123 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, 123)
	}
}

func TestUnmarshalWithZeroFillAndFloat(t *testing.T) {
	line := "hello     00123.45"

	type TestStruct struct {
		Field1 string  `positional:"10"`
		Field2 float64 `positional:"5,zerofill"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != "hello" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, "hello")
	}

	if test.Field2 != 123.45 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, 123.45)
	}
}

func TestUnmarshalWithZeroFillAndFloatAndLeftPad(t *testing.T) {
	line := "hello     00123.45"

	type TestStruct struct {
		Field1 string  `positional:"10"`
		Field2 float64 `positional:"5,zerofill,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != "hello" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, "hello")
	}

	if test.Field2 != 123.45 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, 123.45)
	}
}

func TestUnmarshalWithZeroFillAndFloatAndLeftPadAndShortLine(t *testing.T) {
	line := "hello     00123.45"

	type TestStruct struct {
		Field1 string  `positional:"10"`
		Field2 float64 `positional:"10,zerofill,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err == nil {
		t.Errorf("Expected error for short line, but got none")
	}
}

func TestUnmarshalWithBoolean(t *testing.T) {
	line := "hello     true"

	type TestStruct struct {
		Field1 string `positional:"10"`
		Field2 bool   `positional:"5,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err == nil {
		t.Errorf("Expected error for invalid boolean, but got none")
	}
}

func TestUnmarshalWithBooleanTrue(t *testing.T) {
	line := "hello     true"

	type TestStruct struct {
		Field1 string `positional:"10"`
		Field2 bool   `positional:"4,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != "hello" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, "hello")
	}

	if test.Field2 != true {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, true)
	}
}

func TestUnmarshalWithBooleanFalse(t *testing.T) {
	line := "hello     false"

	type TestStruct struct {
		Field1 string `positional:"10"`
		Field2 bool   `positional:"5,leftpad"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != "hello" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, "hello")
	}

	if test.Field2 != false {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, false)
	}
}

func TestUnmarshalWithAllVariations(t *testing.T) {

	line := "Test       25.77true  25475Test of value   3.54false 0047132546       "

	type TestStruct struct {
		Field1 string  `positional:"10"`
		Field2 float64 `positional:"6,leftpad"`
		Field3 bool    `positional:"5"`
		Field4 int     `positional:"6,leftpad"`
		Field5 string  `positional:"15"`
		Field6 float64 `positional:"5,leftpad"`
		Field7 bool    `positional:"5"`
		Field8 int     `positional:"5,leftpad"`
		Field9 int     `positional:"13"`
	}

	var test TestStruct

	err := positional_line.Unmarshal(line, &test)

	if err != nil {
		t.Errorf("Unmarshal(%v) raised error %v", line, err)
	}

	if test.Field1 != "Test" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field1, "Test")
	}

	if test.Field2 != 25.77 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field2, 25.77)
	}

	if test.Field3 != true {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field3, true)
	}

	if test.Field4 != 25475 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field4, 25475)
	}

	if test.Field5 != "Test of value" {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field5, "Test of value")
	}

	if test.Field6 != 3.54 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field6, 3.54)
	}

	if test.Field7 != false {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field7, false)
	}

	if test.Field8 != 47 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field8, 47)
	}

	if test.Field9 != 132546 {
		t.Errorf("Unmarshal(%v) = %v, expected %v", line, test.Field9, 132546)
	}
}
