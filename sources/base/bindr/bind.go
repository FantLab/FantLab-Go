package bindr

import (
	"errors"
	"reflect"
	"strconv"
)

var ErrBadOutput = errors.New("bindr: output should be a pointer to struct value")

func BindStruct(output interface{}, valueGetter func(f reflect.StructField) string) error {
	val := reflect.ValueOf(output)

	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = reflect.Indirect(val)

		if val.Kind() == reflect.Struct {
			typ := val.Type()

			for i := 0; i < val.NumField(); i++ {
				var s string
				if valueGetter != nil {
					s = valueGetter(typ.Field(i))
				}
				setFieldValue(s, val.Field(i))
			}

			return nil
		}
	}

	return ErrBadOutput
}

func setFieldValue(s string, output reflect.Value) {
	output = reflect.Indirect(output)

	if "" == s || !output.CanSet() {
		return
	}

	switch output.Kind() {
	case reflect.String:
		output.SetString(s)
	case reflect.Bool:
		x, err := strconv.ParseBool(s)
		if err == nil {
			output.SetBool(x)
		}
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:

		x, err := strconv.Atoi(s)
		if err == nil {
			output.SetInt(int64(x))
		}
	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:

		x, err := strconv.ParseUint(s, 10, 0)
		if err == nil {
			output.SetUint(x)
		}
	}
}
