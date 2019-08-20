package scanr

import (
	"fantlab/tt"
	"reflect"
	"testing"
)

// *******************************************************

type _testColumn string

func (column _testColumn) Name() string {
	return string(column)
}

func (column _testColumn) Get(value reflect.Value) reflect.Value {
	return value
}

// *******************************************************

type _testRows struct {
	values  [][]interface{}
	columns []Column
}

func (rows *_testRows) AltNameTag() string {
	return "altname"
}

func (rows *_testRows) IterateUsing(fn RowFunc) error {
	for _, values := range rows.values {
		err := fn(rows.columns, values)

		if err != nil {
			return err
		}
	}

	return nil
}

// *******************************************************

func Test_Scan(t *testing.T) {
	t.Run("negative_output_nil", func(t *testing.T) {
		rows := &_testRows{}

		var x *uint8

		err := Scan(x, rows)

		tt.Assert(t, err == ErrIsNil)
	})

	t.Run("negative_output_not_a_ptr", func(t *testing.T) {
		rows := &_testRows{}

		var x uint8

		err := Scan(x, rows)

		tt.Assert(t, err == ErrNotAPtr)
	})

	t.Run("negative_output_not_a_struct_slice", func(t *testing.T) {
		rows := &_testRows{}

		var x []int

		err := Scan(&x, rows)

		tt.Assert(t, err == ErrNotAStruct)
	})

	t.Run("positive_single_value_1", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{{1}},
			columns: []Column{
				_testColumn(""),
			},
		}

		var x uint8

		err := Scan(&x, rows)

		tt.Assert(t, err == nil)
		tt.Assert(t, x == 1)
	})

	t.Run("positive_single_value_2", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{{"hello"}},
			columns: []Column{
				_testColumn(""),
			},
		}

		var x string

		err := Scan(&x, rows)

		tt.Assert(t, err == nil)
		tt.Assert(t, x == "hello")
	})

	t.Run("negative_single_value_multi_columns", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{{"hello", "world"}},
			columns: []Column{
				_testColumn(""),
				_testColumn(""),
			},
		}

		var x string

		err := Scan(&x, rows)

		tt.Assert(t, err == ErrMultiColumns)
		tt.Assert(t, x == "")
	})

	t.Run("negative_single_value_no_rows", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{},
			columns: []Column{
				_testColumn(""),
			},
		}

		var x string

		err := Scan(&x, rows)

		tt.Assert(t, err == ErrNoRows)
		tt.Assert(t, x == "")
	})

	t.Run("negative_single_value_multi_rows", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{{"hello"}, {"world"}},
			columns: []Column{
				_testColumn(""),
			},
		}

		var x string

		err := Scan(&x, rows)

		tt.Assert(t, err == ErrMultiRows)
		tt.Assert(t, x == "")
	})

	t.Run("positive_single_struct_field_name", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{{"a", "b"}},
			columns: []Column{
				_testColumn("FirstName"),
				_testColumn("LastName"),
			},
		}

		var x struct {
			FirstName string
			LastName  string
		}

		err := Scan(&x, rows)

		tt.Assert(t, err == nil)
		tt.Assert(t, x.FirstName == "a")
		tt.Assert(t, x.LastName == "b")
	})

	t.Run("positive_single_struct_alt_name", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{{"a", "b"}},
			columns: []Column{
				_testColumn("first_name"),
				_testColumn("last_name"),
			},
		}

		var x struct {
			FirstName string `altname:"first_name"`
			LastName  string `altname:"last_name"`
		}

		err := Scan(&x, rows)

		tt.Assert(t, err == nil)
		tt.Assert(t, x.FirstName == "a")
		tt.Assert(t, x.LastName == "b")
	})

	t.Run("negative_single_struct_no_rows", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{},
			columns: []Column{
				_testColumn("first_name"),
				_testColumn("last_name"),
			},
		}

		var x struct {
			FirstName string `altname:"first_name"`
			LastName  string `altname:"last_name"`
		}

		err := Scan(&x, rows)

		tt.Assert(t, err == ErrNoRows)
		tt.Assert(t, x.FirstName == "")
		tt.Assert(t, x.LastName == "")
	})

	t.Run("negative_single_value_multi_rows", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{{"a", "b"}, {"c", "d"}},
			columns: []Column{
				_testColumn("first_name"),
				_testColumn("last_name"),
			},
		}

		var x struct {
			FirstName string `altname:"first_name"`
			LastName  string `altname:"last_name"`
		}

		err := Scan(&x, rows)

		tt.Assert(t, err == ErrMultiRows)
		tt.Assert(t, x.FirstName == "")
		tt.Assert(t, x.LastName == "")
	})

	t.Run("positive_slice_alt_name", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{{"a", "b"}, {"c", "d"}},
			columns: []Column{
				_testColumn("first_name"),
				_testColumn("last_name"),
			},
		}

		var x []struct {
			FirstName string `altname:"first_name"`
			LastName  string `altname:"last_name"`
		}

		err := Scan(&x, rows)

		tt.Assert(t, err == nil)
		tt.Assert(t, x[0].FirstName == "a")
		tt.Assert(t, x[0].LastName == "b")
		tt.Assert(t, x[1].FirstName == "c")
		tt.Assert(t, x[1].LastName == "d")
	})

	t.Run("positive_slice_mix_names", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{{"a", "b"}, {"c", "d"}},
			columns: []Column{
				_testColumn("FirstName"),
				_testColumn("last_name"),
			},
		}

		var x []struct {
			FirstName string
			LastName  string `altname:"last_name"`
		}

		err := Scan(&x, rows)

		tt.Assert(t, err == nil)
		tt.Assert(t, x[0].FirstName == "a")
		tt.Assert(t, x[0].LastName == "b")
		tt.Assert(t, x[1].FirstName == "c")
		tt.Assert(t, x[1].LastName == "d")
	})

	t.Run("positive_complex_slice", func(t *testing.T) {
		rows := &_testRows{
			values: [][]interface{}{
				{"a", "b", 1, 2, true, 2.8},
				{"c", "d", 3, 4, false, 3.2},
			},
			columns: []Column{
				_testColumn("first_name"),
				_testColumn("last_name"),
				_testColumn("id1"),
				_testColumn("id2"),
				_testColumn("is_closed"),
				_testColumn("coef"),
			},
		}

		type testData struct {
			FirstName string  `altname:"first_name"`
			LastName  string  `altname:"last_name"`
			Id1       int     `altname:"id1"`
			Id2       uint8   `altname:"id2"`
			IsClosed  bool    `altname:"is_closed"`
			Coef      float64 `altname:"coef"`
		}

		var x []testData

		err := Scan(&x, rows)

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, x[0], testData{
			FirstName: "a",
			LastName:  "b",
			Id1:       1,
			Id2:       2,
			IsClosed:  true,
			Coef:      2.8,
		})
		tt.AssertDeepEqual(t, x[1], testData{
			FirstName: "c",
			LastName:  "d",
			Id1:       3,
			Id2:       4,
			IsClosed:  false,
			Coef:      3.2,
		})
	})
}
