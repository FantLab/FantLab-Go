package scheme

import (
	"reflect"
	"strings"
)

type BuilderConfig struct {
	GetComment           func(reflect.Type, string) string
	IsValidField         func(reflect.StructField) bool
	GetFieldName         func(reflect.StructTag) string
	CustomStructStringer func(reflect.Type) string
}

func NewBuilder(cfg *BuilderConfig) *Builder {
	if cfg == nil {
		cfg = new(BuilderConfig)
	}
	if cfg.GetFieldName == nil {
		cfg.GetFieldName = func(tag reflect.StructTag) string {
			return strings.Split(tag.Get("json"), ",")[0]
		}
	}
	return &Builder{cfg: cfg}
}

type Builder struct {
	cfg *BuilderConfig
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
		if b.cfg.CustomStructStringer != nil {
			if s := b.cfg.CustomStructStringer(t); s != "" {
				ls.current.builder.WriteString(s)
				return
			}
		}

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
				if b.cfg.IsValidField != nil && !b.cfg.IsValidField(f) {
					continue
				}

				ls.new(d + 1)
				ls.current.builder.WriteString(b.cfg.GetFieldName(f.Tag) + ": ")
				if b.cfg.GetComment != nil {
					ls.current.comment = b.cfg.GetComment(t, f.Name)
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
