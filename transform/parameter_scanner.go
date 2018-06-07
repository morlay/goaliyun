package transform

import (
	"encoding"
	"errors"
	"fmt"
	"go/ast"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func NewParameterScanner() *ParameterScanner {
	return &ParameterScanner{
		params:     url.Values{},
		pathParams: url.Values{},
	}
}

type ParameterScanner struct {
	walker     PathWalker
	params     url.Values
	pathParams url.Values
}

func (s *ParameterScanner) SetParam(key string, value string, position string) {
	if value != "" {
		switch position {
		case "Query":
			s.params.Set(key, value)
		case "Path":
			s.pathParams.Set(key, value)
		}
	}
}

func (s *ParameterScanner) Scan(rv reflect.Value, position string) error {
	v := rv.Interface()
	if rv.Kind() != reflect.Ptr {
		if rv.CanAddr() {
			v = rv.Addr().Interface()
		}
	} else if rv.IsNil() {
		return nil
	}

	if stringer, ok := v.(fmt.Stringer); ok {
		s.SetParam(s.walker.Current(), stringer.String(), position)
		return nil
	}

	if textMarshaler, ok := v.(encoding.TextMarshaler); ok {
		text, err := textMarshaler.MarshalText()
		if err != nil {
			return err
		}

		s.SetParam(s.walker.Current(), string(text), position)
		return nil
	}

	rv = reflect.Indirect(rv)

	switch rv.Kind() {
	case reflect.Struct:
		tpe := rv.Type()

		for i := 0; i < tpe.NumField(); i++ {
			field := tpe.Field(i)
			if !ast.IsExported(field.Name) {
				continue
			}

			name := field.Name

			if jsonTag, exists := field.Tag.Lookup("json"); exists {
				name, _ = tagAngFlags(jsonTag)
			}

			if nameTag, exists := field.Tag.Lookup("name"); exists {
				name, _ = tagAngFlags(nameTag)
			}

			if positionValue, exists := field.Tag.Lookup("position"); exists {
				position = positionValue
			}

			if (name != "-") || field.Anonymous {
				if !field.Anonymous {
					s.walker.Enter(name)
				}

				if err := s.Scan(rv.Field(i), position); err != nil {
					return err
				}

				if !field.Anonymous {
					s.walker.Exit()
				}
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			s.walker.Enter(i)
			if err := s.Scan(rv.Index(i), position); err != nil {
				return err
			}
			s.walker.Exit()
		}
	default:
		v := rv.Interface()

		if isEmptyValue(rv) {
			return nil
		}

		switch rv.Kind() {
		case reflect.String:
			s.SetParam(s.walker.Current(), v.(string), position)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
			s.SetParam(s.walker.Current(), fmt.Sprintf("%d", v), position)
		case reflect.Bool:
			s.SetParam(s.walker.Current(), strconv.FormatBool(v.(bool)), position)
		case reflect.Float32:
			s.SetParam(s.walker.Current(), strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32), position)
		case reflect.Float64:
			s.SetParam(s.walker.Current(), strconv.FormatFloat(v.(float64), 'f', -1, 64), position)
		default:
			return ErrorUnSupportedType
		}
	}
	return nil
}

var (
	ErrorUnSupportedType = errors.New("unsupported type")
)

func (s *ParameterScanner) Params() url.Values {
	return s.params
}

func (s *ParameterScanner) PathParams() url.Values {
	return s.pathParams
}

func tagAngFlags(tag string) (string, map[string]bool) {
	values := strings.Split(tag, ",")

	name := values[0]
	flags := map[string]bool{}

	if len(values) > 1 {
		for _, v := range values[1:] {
			flags[v] = true
		}
	}

	return name, flags
}

func isEmptyValue(rv reflect.Value) bool {
	if rv.IsValid() && rv.CanInterface() {
		if zeroChecker, ok := rv.Interface().(interface{ IsZero() bool }); ok {
			return zeroChecker.IsZero()
		}
	}

	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return rv.IsNil()
	}
	return false
}
