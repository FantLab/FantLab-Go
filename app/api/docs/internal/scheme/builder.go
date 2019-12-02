package scheme

import (
	"reflect"
	"strings"
)

func NewBuilder(
	getComment func(t reflect.Type, fieldName string) string,
	isValidField func(f reflect.StructField) bool) *Builder {

	return &Builder{
		getComment:   getComment,
		isValidField: isValidField,
	}
}

type Builder struct {
	getComment   func(t reflect.Type, fieldName string) string
	isValidField func(f reflect.StructField) bool
}

func (b *Builder) Make(t reflect.Type, prefix, postfix string) string {
	if t == nil {
		return ""
	}

	ls := new(lines)
	ls.new(0)
	b.walkType(ls, nil, unptr(t))
	ls.new(0)
	return ls.join(prefix, postfix)
}

func (b *Builder) walkType(ls *lines, superTypes []string, t reflect.Type) {
	switch t.Kind() {
	case reflect.Struct:
		for _, st := range superTypes {
			if st == t.String() {
				ls.current.builder.WriteString("...")
				return
			}
		}

		ls.current.builder.WriteRune('{')
		{
			d := len(superTypes)

			hasFields := false

			for i := 0; i < t.NumField(); i++ {
				f := t.Field(i)
				if b.isValidField != nil && !b.isValidField(f) {
					continue
				}

				ls.new(d + 1)
				ls.current.builder.WriteString(strings.Split(f.Tag.Get("json"), ",")[0] + ": ")
				if b.getComment != nil {
					ls.current.comment = b.getComment(t, f.Name)
				}

				b.walkType(ls, append(superTypes, t.String()), unptr(f.Type))

				hasFields = true
			}

			if hasFields {
				ls.new(d)
			}
		}
		ls.current.builder.WriteRune('}')
	case reflect.Slice:
		ls.current.builder.WriteRune('[')
		b.walkType(ls, superTypes, unptr(t.Elem()))
		ls.current.builder.WriteRune(']')
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

		ls.current.builder.WriteString(t.Kind().String())
	}
}
