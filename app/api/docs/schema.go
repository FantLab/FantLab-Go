package docs

import (
	"reflect"
	"strings"
)

func getScheme(t reflect.Type, indent string) string {
	sb := new(strings.Builder)
	walkType(sb, indent, nil, unptr(t))
	return sb.String()
}

func unptr(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

func walkType(sb *strings.Builder, indent string, superTypes []string, t reflect.Type) {
	switch t.Kind() {
	case reflect.Struct:
		for _, st := range superTypes {
			if st == t.String() {
				sb.WriteString("...")
				return
			}
		}

		sb.WriteRune('{')
		{
			d := len(superTypes)

			hasFields := false

			for i := 0; i < t.NumField(); i++ {
				f := t.Field(i)
				if strings.HasPrefix(f.Name, "XXX") {
					continue
				}

				sb.WriteRune('\n')
				for i := 0; i <= d; i++ {
					sb.WriteString(indent)
				}
				sb.WriteString(strings.Split(f.Tag.Get("json"), ",")[0])
				sb.WriteRune(':')
				sb.WriteRune(' ')

				walkType(sb, indent, append(superTypes, t.String()), unptr(f.Type))

				hasFields = true
			}

			if hasFields {
				sb.WriteRune('\n')
				for i := 0; i < d; i++ {
					sb.WriteString(indent)
				}
			}
		}
		sb.WriteRune('}')
	case reflect.Slice:
		sb.WriteRune('[')
		{
			walkType(sb, indent, superTypes, unptr(t.Elem()))
		}
		sb.WriteRune(']')
	case
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.String:

		sb.WriteString(t.Kind().String())
	}
}
