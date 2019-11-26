package docs

import (
	"reflect"
	"strings"
)

type schemeBuilder struct {
	indent       string
	getComment   func(t reflect.Type, fieldName string) string
	isValidField func(f reflect.StructField) bool
}

func (b *schemeBuilder) make(t reflect.Type) string {
	sb := new(strings.Builder)
	b.walkType(sb, nil, unptr(t))
	return sb.String()
}

func (b *schemeBuilder) walkType(sb *strings.Builder, superTypes []string, t reflect.Type) {
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
				if b.isValidField != nil && !b.isValidField(f) {
					continue
				}

				sb.WriteRune('\n')
				for i := 0; i <= d; i++ {
					sb.WriteString(b.indent)
				}
				sb.WriteString(strings.Split(f.Tag.Get("json"), ",")[0] + ": ")

				b.walkType(sb, append(superTypes, t.String()), unptr(f.Type))

				if b.getComment != nil {
					comment := b.getComment(t, f.Name)
					if "" != comment {
						sb.WriteString(" # " + comment)
					}
				}

				hasFields = true
			}

			if hasFields {
				sb.WriteRune('\n')
				for i := 0; i < d; i++ {
					sb.WriteString(b.indent)
				}
			}
		}
		sb.WriteRune('}')
	case reflect.Slice:
		sb.WriteRune('[')
		b.walkType(sb, superTypes, unptr(t.Elem()))
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
