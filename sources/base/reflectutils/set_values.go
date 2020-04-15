package reflectutils

import (
	"reflect"
	"strconv"
)

func SetStructValues(output interface{}, tagName string, valueGetter func(string) string) {
	if valueGetter == nil {
		return
	}

	val := reflect.Indirect(reflect.ValueOf(output))
	if val.Kind() != reflect.Struct {
		return
	}

	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		s := valueGetter(typ.Field(i).Tag.Get(tagName))
		f := reflect.Indirect(val.Field(i))
		if f.CanSet() {
			setValueToField(s, f)
		}
	}
}

func setValueToField(value string, field reflect.Value) {
	switch field.Kind() {
	case reflect.String:
		if value != "" {
			field.SetString(value)
		}
	case reflect.Bool:
		x, err := strconv.ParseBool(value)
		if err == nil {
			field.SetBool(x)
		}
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:

		x, err := strconv.Atoi(value)
		if err == nil {
			field.SetInt(int64(x))
		}
	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:

		x, err := strconv.ParseUint(value, 10, 0)
		if err == nil {
			field.SetUint(x)
		}
	}
}
