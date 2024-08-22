package positional_line

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/vert-capital/positional_line/pad"
)

func ParseTags(t reflect.Type) (TagCollection, error) {
	var tags []Tag

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		ftag := field.Tag.Get(tagName)

		if ftag == "" {
			continue
		}

		opts := strings.Split(ftag, ",")

		size, err := strconv.Atoi(opts[0])

		if err != nil {
			return TagCollection{}, ErrInvalidSize
		}

		zerofill := false
		leftpad := false
		nofloat := false

		modifiers := opts[1:]

		for _, m := range modifiers {
			if m == "zerofill" {
				zerofill = true
			}

			if m == "leftpad" {
				leftpad = true
			}

			if m == "nofloat" {
				nofloat = true
			}
		}

		t := Tag{
			Name:     field.Name,
			Size:     size,
			LeftPad:  leftpad,
			ZeroFill: zerofill,
			NoFloat:  nofloat,
		}

		tags = append(tags, t)
	}

	line := TagCollection{
		Name: t.Name(),
		Tags: tags,
	}

	return line, nil
}

func UnparseValue(rv reflect.Value, line TagCollection, content string) error {
	var err error
	t := rv.Type()
	start := 0

	for i := 0; i < rv.NumField(); i++ {
		field := t.Field(i)
		value := rv.Field(i)

		tg := Tags(line)[field.Name]

		if tg == (Tag{}) {
			continue
		}

		end := start + tg.Size

		fieldContent := content[start:end]

		err = Unconvert(value, tg, fieldContent)

		if err != nil {
			return err
		}

		start = end
	}

	return nil
}

func Unconvert(v reflect.Value, t Tag, content string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(strings.TrimSpace(content))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(strings.TrimSpace(content), 10, 64)

		if err != nil {
			return err
		}

		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := strconv.ParseUint(strings.TrimSpace(content), 10, 64)

		if err != nil {
			return err
		}

		v.SetUint(i)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(strings.TrimSpace(content), 64)

		if err != nil {
			return err
		}

		v.SetFloat(f)
	case reflect.Bool:
		b, err := strconv.ParseBool(strings.TrimSpace(content))

		if err != nil {
			return err
		}

		v.SetBool(b)
	}

	return nil
}

func ParseValue(rv reflect.Value, line TagCollection) (string, error) {
	var err error
	var content strings.Builder
	t := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		field := t.Field(i)
		value := rv.Field(i)

		tg := Tags(line)[field.Name]

		if tg == (Tag{}) {
			continue
		}

		fieldContent := Convert(value, tg)

		var sep string
		if tg.ZeroFill {
			sep = "0"
		} else {
			sep = " "
		}

		var fline string
		if tg.LeftPad {
			fline, err = pad.Left(fieldContent, tg.Size, sep)
		} else {
			fline, err = pad.Right(fieldContent, tg.Size, sep)
		}

		if err != nil {
			return "", err
		}

		content.WriteString(fline)
	}

	return content.String(), nil
}

func Convert(v reflect.Value, t Tag) string {
	var content string

	switch v.Kind() {
	case reflect.String:
		content = v.Interface().(string)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		content = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		content = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		value := fmt.Sprintf("%.2f", v.Float())

		if t.NoFloat {
			content = strings.Replace(value, ".", "", 1)
		} else {
			content = value
		}
	case reflect.Bool:
		value := v.Interface().(bool)

		if value {
			content = "1"
		} else {
			content = "0"
		}
	}

	return content
}

func Tags(l TagCollection) map[string]Tag {
	tags := make(map[string]Tag)

	for _, t := range l.Tags {
		tags[t.Name] = t
	}

	return tags
}
