package positional_line_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vert-capital/positional_line"
)

func TestParseTags(t *testing.T) {

	type TestParseTagsStruct struct {
		Field1 string  `positional:"10,zerofill,leftpad"`
		Field2 float64 `positional:"7,nofloat"`
		Field3 bool
	}

	testStruct := TestParseTagsStruct{
		Field1: "123",
		Field2: 456,
		Field3: true,
	}

	tags, err := positional_line.ParseTags(reflect.TypeOf(testStruct))
	if err != nil {
		t.Fatalf("parseTags returned error: %v", err)
	}

	if len(tags.Tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(tags.Tags))
	}

	if len(tags.Tags) == 0 {
		t.Fatalf("No tags found")
	}

	if tags.Tags[0].Name != "Field1" || tags.Tags[0].Size != 10 || !tags.Tags[0].ZeroFill || !tags.Tags[0].LeftPad || tags.Tags[0].NoFloat {
		t.Errorf("Incorrect tag values for Field1: %+v", tags.Tags[0])
	}

	if tags.Tags[1].Name != "Field2" || tags.Tags[1].Size != 7 || tags.Tags[1].ZeroFill || tags.Tags[1].LeftPad || !tags.Tags[1].NoFloat {
		t.Errorf("Incorrect tag values for Field2: %+v", tags.Tags[1])
	}
}
func TestParseValue(t *testing.T) {

	type TestStruct struct {
		Field1 string  `positional:"10,zerofill,leftpad"`
		Field2 float64 `positional:"7,nofloat,leftpad"`
		Field3 bool    `positional:"7,leftpad"`
	}

	tags, err := positional_line.ParseTags(reflect.TypeOf(TestStruct{}))
	if err != nil {
		t.Fatalf("parseTags returned error: %v", err)
	}

	tests := []struct {
		input    TestStruct
		expected string
	}{
		{TestStruct{"hello", 123.70, true}, "00000hello  12370      1"},
		{TestStruct{"world", 456, false}, "00000world  45600      0"},
	}

	for _, test := range tests {
		result, err := positional_line.ParseValue(reflect.ValueOf(test.input), tags)
		if err != nil {
			t.Errorf("parseValue returned error: %v", err)
		}
		if result != test.expected {
			t.Errorf("ParseValue(%v) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestTestParseValueEspecificPositional(t *testing.T) {
	type TestStruct1 struct {
		FieldInt    int     `positional:"10"`
		FieldFloat  float64 `positional:"10"`
		FieldString string  `positional:"10"`
		FieldBool   bool    `positional:"1"`
	}

	testStruct := TestStruct1{
		FieldInt:    123,
		FieldFloat:  123.70,
		FieldString: "hello",
		FieldBool:   true,
	}

	tags, err := positional_line.ParseTags(reflect.TypeOf(testStruct))

	assert.Nil(t, err)

	result, err := positional_line.ParseValue(reflect.ValueOf(testStruct), tags)
	assert.Nil(t, err)
	assert.Equal(t, "123       123.70    hello     1", result)
}

func TestTestParseValueEspecificLeftPad(t *testing.T) {
	type TestStruct1 struct {
		FieldInt    int     `positional:"10,leftpad"`
		FieldFloat  float64 `positional:"10,leftpad"`
		FieldString string  `positional:"10,leftpad"`
		FieldBool   bool    `positional:"2,leftpad"`
	}

	testStruct := TestStruct1{
		FieldInt:    123,
		FieldFloat:  123.70,
		FieldString: "hello",
		FieldBool:   true,
	}

	tags, err := positional_line.ParseTags(reflect.TypeOf(testStruct))

	assert.Nil(t, err)

	result, err := positional_line.ParseValue(reflect.ValueOf(testStruct), tags)
	assert.Nil(t, err)
	assert.Equal(t, "       123    123.70     hello 1", result)
}

func TestTestParseValueEspecificZeroFill(t *testing.T) {
	type TestStruct1 struct {
		FieldInt    int     `positional:"10,zerofill"`
		FieldFloat  float64 `positional:"10,zerofill"`
		FieldString string  `positional:"10,zerofill"`
		FieldBool   bool    `positional:"1,zerofill"`
	}

	testStruct := TestStruct1{
		FieldInt:    123,
		FieldFloat:  123.70,
		FieldString: "hello",
		FieldBool:   true,
	}

	tags, err := positional_line.ParseTags(reflect.TypeOf(testStruct))

	assert.Nil(t, err)

	result, err := positional_line.ParseValue(reflect.ValueOf(testStruct), tags)
	assert.Nil(t, err)
	assert.Equal(t, "1230000000123.700000hello000001", result)
}

func TestTestParseValueEspecificZeroFillLeftPad(t *testing.T) {
	type TestStruct1 struct {
		FieldInt    int     `positional:"10,zerofill,leftpad"`
		FieldFloat  float64 `positional:"10,zerofill,leftpad"`
		FieldString string  `positional:"10,zerofill,leftpad"`
		FieldBool   bool    `positional:"2,zerofill,leftpad"`
	}

	testStruct := TestStruct1{
		FieldInt:    123,
		FieldFloat:  123.70,
		FieldString: "hello",
		FieldBool:   true,
	}

	tags, err := positional_line.ParseTags(reflect.TypeOf(testStruct))

	assert.Nil(t, err)

	result, err := positional_line.ParseValue(reflect.ValueOf(testStruct), tags)
	assert.Nil(t, err)
	assert.Equal(t, "00000001230000123.7000000hello01", result)
}

func TestTestParseValueEspecificNoFloat(t *testing.T) {
	type TestStruct1 struct {
		FieldInt    int     `positional:"10,nofloat"`
		FieldFloat  float64 `positional:"10,nofloat"`
		FieldString string  `positional:"10,nofloat"`
		FieldBool   bool    `positional:"1,nofloat"`
	}

	testStruct := TestStruct1{
		FieldInt:    123,
		FieldFloat:  123.70,
		FieldString: "hello",
		FieldBool:   true,
	}

	tags, err := positional_line.ParseTags(reflect.TypeOf(testStruct))

	assert.Nil(t, err)

	result, err := positional_line.ParseValue(reflect.ValueOf(testStruct), tags)
	assert.Nil(t, err)
	assert.Equal(t, "123       12370     hello     1", result)
}

func TestTestParseValueEspecificNoFloatLeftPad(t *testing.T) {
	type TestStruct1 struct {
		FieldInt    int     `positional:"10,nofloat,leftpad"`
		FieldFloat  float64 `positional:"10,nofloat,leftpad"`
		FieldString string  `positional:"10,nofloat,leftpad"`
		FieldBool   bool    `positional:"2,nofloat,leftpad"`
	}

	testStruct := TestStruct1{
		FieldInt:    123,
		FieldFloat:  123.70,
		FieldString: "hello",
		FieldBool:   true,
	}

	tags, err := positional_line.ParseTags(reflect.TypeOf(testStruct))

	assert.Nil(t, err)

	result, err := positional_line.ParseValue(reflect.ValueOf(testStruct), tags)
	assert.Nil(t, err)
	assert.Equal(t, "       123     12370     hello 1", result)
}

func TestTestParseValueEspecificNoFloatLeftPadZerofill(t *testing.T) {
	type TestStruct1 struct {
		FieldInt    int     `positional:"10,nofloat,leftpad,zerofill"`
		FieldFloat  float64 `positional:"10,nofloat,leftpad,zerofill"`
		FieldString string  `positional:"10,nofloat,leftpad,zerofill"`
		FieldBool   bool    `positional:"2,nofloat,leftpad,zerofill"`
	}

	testStruct := TestStruct1{
		FieldInt:    123,
		FieldFloat:  123.70,
		FieldString: "hello",
		FieldBool:   true,
	}

	tags, err := positional_line.ParseTags(reflect.TypeOf(testStruct))

	assert.Nil(t, err)

	result, err := positional_line.ParseValue(reflect.ValueOf(testStruct), tags)
	assert.Nil(t, err)
	assert.Equal(t, "0000000123000001237000000hello01", result)
}

func TestConvert(t *testing.T) {
	assert.Equal(t,
		positional_line.Convert(reflect.ValueOf(123), positional_line.Tag{}),
		"123",
	)

	assert.Equal(t,
		positional_line.Convert(reflect.ValueOf(123), positional_line.Tag{}),
		"123",
	)

	assert.Equal(t,
		positional_line.Convert(reflect.ValueOf(123.70), positional_line.Tag{}),
		"123.70",
	)

	assert.Equal(t,
		positional_line.Convert(reflect.ValueOf(123.70), positional_line.Tag{NoFloat: true}),
		"12370",
	)

	assert.Equal(t,
		positional_line.Convert(reflect.ValueOf(false), positional_line.Tag{}),
		"0",
	)

	assert.Equal(t,
		positional_line.Convert(reflect.ValueOf(true), positional_line.Tag{}),
		"1",
	)
}

func TestTags(t *testing.T) {
	tests := []struct {
		input    positional_line.TagCollection
		expected map[string]positional_line.Tag
	}{
		{
			positional_line.TagCollection{
				Tags: []positional_line.Tag{
					{Name: "Field1", Size: 10, ZeroFill: true, LeftPad: true},
					{Name: "Field2", Size: 5, NoFloat: true},
				},
			},
			map[string]positional_line.Tag{
				"Field1": {Name: "Field1", Size: 10, ZeroFill: true, LeftPad: true},
				"Field2": {Name: "Field2", Size: 5, NoFloat: true},
			},
		},
		{
			positional_line.TagCollection{
				Tags: []positional_line.Tag{
					{Name: "Field3", Size: 15, ZeroFill: false, LeftPad: false},
				},
			},
			map[string]positional_line.Tag{
				"Field3": {Name: "Field3", Size: 15, ZeroFill: false, LeftPad: false},
			},
		},
	}

	for _, test := range tests {
		result := positional_line.Tags(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("tags(%+v) = %+v; want %+v", test.input, result, test.expected)
		}
	}
}
