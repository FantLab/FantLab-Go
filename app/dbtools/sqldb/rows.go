package sqldb

import (
	"database/sql"
	"fantlab/dbtools/scanr"
	"reflect"
	"strings"
)

// *******************************************************

type sqlRows struct {
	data           *sql.Rows
	err            error
	allowNullTypes bool
}

func (rows sqlRows) Error() error {
	return rows.err
}

func (rows sqlRows) Scan(output interface{}) error {
	if rows.err != nil {
		return rows.err
	}

	return scanr.Scan(output, rows)
}

func (rows sqlRows) AltNameTag() string {
	return "db"
}

func (rows sqlRows) IterateUsing(fn scanr.RowFunc) error {
	if rows.err != nil {
		return rows.err
	}

	defer rows.data.Close()

	columnTypes, err := rows.data.ColumnTypes()

	if err != nil {
		return err
	}

	values, columns := getColumnData(columnTypes, rows.allowNullTypes)

	for rows.data.Next() {
		err = rows.data.Scan(values...)

		if err != nil {
			return err
		}

		err = fn(columns, values)

		if err != nil {
			return err
		}
	}

	return rows.data.Err()
}

// *******************************************************

type sqlColumn struct {
	name                string
	takeNonNullSubField bool
}

func (column *sqlColumn) Name() string {
	return column.name
}

func (column *sqlColumn) Get(value reflect.Value) reflect.Value {
	if column.takeNonNullSubField {
		return value.Elem().Field(0)
	}

	return value.Elem()
}

// *******************************************************

func getColumnData(columnTypes []*sql.ColumnType, allowNullTypes bool) ([]interface{}, []scanr.Column) {
	size := len(columnTypes)

	values := make([]interface{}, size)
	columns := make([]scanr.Column, size)

	for i, columnType := range columnTypes {
		values[i] = reflect.New(columnType.ScanType()).Interface()

		isNullable := strings.HasPrefix(columnType.ScanType().Name(), "Null")

		columns[i] = &sqlColumn{
			name:                columnType.Name(),
			takeNonNullSubField: isNullable && !allowNullTypes,
		}
	}

	return values, columns
}
