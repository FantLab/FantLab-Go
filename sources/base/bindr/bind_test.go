package bindr

import (
	"fantlab/base/assert"
	"reflect"
	"testing"
)

func Test_setFieldValue(t *testing.T) {
	t.Run("set_string", func(t *testing.T) {
		var output string
		setFieldValue("foo", reflect.ValueOf(&output))
		assert.True(t, output == "foo")
	})

	t.Run("set_bool", func(t *testing.T) {
		var output bool
		setFieldValue("true", reflect.ValueOf(&output))
		assert.True(t, output == true)
	})

	t.Run("set_int", func(t *testing.T) {
		var output int
		setFieldValue("100", reflect.ValueOf(&output))
		assert.True(t, output == 100)
	})

	t.Run("set_uint_positive", func(t *testing.T) {
		var output uint
		setFieldValue("100", reflect.ValueOf(&output))
		assert.True(t, output == 100)
	})

	t.Run("set_uint_negative", func(t *testing.T) {
		var output uint
		setFieldValue("-100", reflect.ValueOf(&output))
		assert.True(t, output == 0)
	})
}

func Test_BindStruct(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		type x struct {
			Id   int
			Name string
		}
		var output x
		err := BindStruct(&output, func(f reflect.StructField) string {
			switch f.Name {
			case "Id":
				return "100"
			case "Name":
				return "user"
			}
			return ""
		})
		assert.True(t, err == nil)
		assert.True(t, output.Id == 100)
		assert.True(t, output.Name == "user")
	})

	t.Run("negative", func(t *testing.T) {
		var output string
		err := BindStruct(output, nil)
		assert.True(t, err == ErrBadOutput)
	})
}
