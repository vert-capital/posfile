package positional_line

import (
	"errors"
	"reflect"
	"strings"
)

const tagName string = "positional"

var (
	// ErrInvalidSize is raised when the size tag are not int
	ErrInvalidSize = errors.New("posline: tag size should be an integer")
)

type TagCollection struct {
	Name string
	Tags []Tag
}

type Tag struct {
	Name     string
	Size     int
	LeftPad  bool
	ZeroFill bool
	NoFloat  bool
}

// Marshal parsers all structs and transform into one string with all lines
func Marshal(v interface{}) (string, error) {
	var lines strings.Builder

	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Struct:
		l, err := marshalStruct(rv)

		if err != nil {
			return "", err
		}

		lines.WriteString(l)
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			l, err := marshalStruct(rv.Index(i))

			if err != nil {
				return "", err
			}

			lines.WriteString(l)

			if i != (rv.Len() - 1) {
				lines.WriteString("\n")
			}
		}
	}

	return lines.String(), nil
}

func marshalStruct(rv reflect.Value) (string, error) {
	var c TagCollection

	t := rv.Type()

	c, err := ParseTags(t)

	if err != nil {
		return "", err
	}

	content, err := ParseValue(rv, c)

	return content, err
}
