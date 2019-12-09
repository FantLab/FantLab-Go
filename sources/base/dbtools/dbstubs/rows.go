package dbstubs

import (
	"fantlab/base/dbtools/scanr"
	"reflect"
)

type StubRows struct {
	Values  [][]interface{}
	Columns []scanr.Column
	Err     error
}

func (rows *StubRows) Error() error {
	return rows.Err
}

func (rows *StubRows) Scan(output interface{}) error {
	if rows.Err != nil {
		return rows.Err
	}

	return scanr.Scan(output, rows)
}

func (rows *StubRows) AltNameTag() string {
	return "db"
}

func (rows *StubRows) IterateUsing(fn scanr.RowFunc) error {
	if rows.Err != nil {
		return rows.Err
	}

	for _, values := range rows.Values {
		err := fn(rows.Columns, values)

		if err != nil {
			return err
		}
	}

	return nil
}

// *******************************************************

type StubColumn string

func (column StubColumn) Name() string {
	return string(column)
}

func (column StubColumn) Get(value reflect.Value) reflect.Value {
	return value
}
