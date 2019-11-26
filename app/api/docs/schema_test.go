package docs

import (
	"fantlab/assert"
	"reflect"
	"strings"
	"testing"
)

func isSchemeEquals(realScheme, expectedScheme string) bool {
	expectedScheme = strings.Trim(strings.ReplaceAll(expectedScheme, "\t", ""), "\n")
	return realScheme == expectedScheme
}

func Test_getScheme(t *testing.T) {
	t.Run("scheme_1", func(t *testing.T) {
		type x struct {
			Id   int    `json:"id,omitempty"`
			Text string `json:"text,omitempty"`
		}

		s := new(schemeBuilder).make(reflect.TypeOf(new(x)))

		assert.True(t, isSchemeEquals(s, `
		{
			id: int
			text: string
		}
		`))
	})

	t.Run("scheme_2", func(t *testing.T) {
		type x struct {
			Id   int `json:"id,omitempty"`
			Text struct {
				UserId int     `json:"user_id,omitempty"`
				Date   float64 `json:"date,omitempty"`
			} `json:"text,omitempty"`
		}

		s := new(schemeBuilder).make(reflect.TypeOf(new(x)))

		assert.True(t, isSchemeEquals(s, `
		{
			id: int
			text: {
				user_id: int
				date: float64
			}
		}
		`))
	})

	t.Run("scheme_recursive", func(t *testing.T) {
		type x struct {
			Id       int `json:"id,omitempty"`
			Children []x `json:"children,omitempty"`
		}

		s := new(schemeBuilder).make(reflect.TypeOf(new(x)))

		assert.True(t, isSchemeEquals(s, `
		{
			id: int
			children: [...]
		}
		`))
	})

	t.Run("scheme_complex", func(t *testing.T) {
		type x struct {
			Id    int `json:"id,omitempty"`
			Items []struct {
				ItemId string `json:"item_id,omitempty"`
				Value  bool   `json:"value,omitempty"`
			} `json:"items,omitempty"`
			Data struct {
				Name string `json:"name,omitempty"`
			} `json:"data,omitempty"`
			Children []x `json:"children,omitempty"`
		}

		s := new(schemeBuilder).make(reflect.TypeOf(new(x)))

		assert.True(t, isSchemeEquals(s, `
		{
			id: int
			items: [{
				item_id: string
				value: bool
			}]
			data: {
				name: string
			}
			children: [...]
		}
		`))
	})
}
